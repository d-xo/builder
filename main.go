package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "env"
	app.Usage = "easily work with dockerized build envionments"

	app.Commands = []cli.Command{
		{
			Name:   "up",
			Usage:  "spin up the project build envionment",
			Action: up,
		},
		{
			Name:   "attach",
			Usage:  "Attach to the project build environment. Will bring envionment up if needed",
			Action: attach,
		},
		{
			Name:   "destroy",
			Usage:  "destroy the project build environment",
			Action: destroy,
		},
		{
			Name:   "clean",
			Usage:  "destroy and rebuild the current environment",
			Action: clean,
		},
	}

	app.Run(os.Args)
}

//
// COMMANDS
//

func up(context *cli.Context) {
	config := loadConfig()
	imageID := buildImage(config)
	startBackgroundContainer(imageID, config.Volumes)
}

func attach(context *cli.Context) {
	docker("exec", "-i", "-t", ContainerName(), "/bin/bash")
}

func destroy(context *cli.Context) {
	docker("rm", "--force", ContainerName())
}

func clean(context *cli.Context) {
	destroy(context)
	up(context)
}

//
// DATA
//

// ConfigPath <- recurse upwards until .workspace.json is found
func ConfigPath() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return findConfig(currentDirectory)
}

// Config <- .workspace.json
type Config struct {
	DockerfileDirectory string            `json:"dockerfileDirectory"`
	Volumes             map[string]string `json:"volumes"`
}

// ContainerName <- hash of ConfigPath()
func ContainerName() string {
	return hash([]byte(ConfigPath()))
}

// ConfigDir <- directory of config file
func ConfigDir() string {
	return filepath.Dir(ConfigPath())
}

//
// CONTROL DOCKER
//

func buildImage(config Config) string {
	stdoutStderr, err := exec.Command("docker", "build", "--quiet", config.DockerfileDirectory).CombinedOutput()
	if err != nil {
		fmt.Println(string(stdoutStderr))
		panic(err)
	}

	imageID := strings.TrimSpace(strings.Split(string(stdoutStderr), ":")[1])

	fmt.Println("built image with ID:", imageID)
	fmt.Println("from Dockerfile in:", config.DockerfileDirectory)

	return imageID
}

func startBackgroundContainer(imageID string, volumes map[string]string) {
	docker("run", "-dti", dockerArgsVolumesString(volumes), "--name", ContainerName(), imageID)
}

//
// CONFIG
//

func loadConfig() Config {
	configFile, err := os.Open(ConfigPath())
	if err != nil {
		panic(err)
	}

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panic(err)
	}

	return normalizeConfig(config)
}

// normalizeConfig: make all host paths absolute
func normalizeConfig(config Config) Config {
	config.DockerfileDirectory = normalizePath(config.DockerfileDirectory)
	config.Volumes = normalizeVolumes(config.Volumes)
	return config
}

func normalizeVolumes(volumes map[string]string) map[string]string {
	normalizedVolumes := make(map[string]string)
	for host, guest := range volumes {
		normalizedVolumes[normalizePath(host)] = guest
	}
	return normalizedVolumes
}

func normalizePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(ConfigDir(), path)
}

func findConfig(directory string) string {
	configFile := filepath.Join(directory, ".workspace.json")

	if pathExists(configFile) {
		return configFile
	}

	if isPathRoot(directory) {
		fmt.Println("No .workspace.json found")
		os.Exit(1)
	}

	return findConfig(filepath.Dir(directory))
}

func dockerArgsVolumesString(volumes map[string]string) string {
	output := ""

	for host, dest := range volumes {
		volumeCommand := []string{output, "-v", normalizePath(host), ":", dest, ""}
		output = strings.Join(volumeCommand, "")
	}

	return output
}

//
// util
//

func docker(args ...string) {
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func hash(bytes []byte) string {
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func isPathRoot(path string) bool {
	if filepath.Dir(path) == filepath.Dir(filepath.Dir(path)) {
		return true
	}
	return false
}
