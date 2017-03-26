package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"path/filepath"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "env"
	app.Usage = "use dockerfiles to define per project working environments"

	app.Commands = []cli.Command{
		{
			Name:   "up",
			Usage:  "spin up a an environment from the dockerfile in the current directory",
			Action: up,
		},
		{
			Name:   "attach",
			Usage:  "attach to the environment for the current directory",
			Action: attach,
		},
		{
			Name:   "destroy",
			Usage:  "destroy the environment for the current directory",
			Action: destroy,
		},
		{
			Name:   "reset",
			Usage:  "reset the environment to the state defined in the dockerfile",
			Action: reset,
		},
	}

	app.Run(os.Args)
}

//
// DATA
//

// Config <- .weatherconfig.json
type Config struct {
	DockerfileDirectory string            `json:"dockerfileDirectory"`
	Volumes             map[string]string `json:"volumes"`
}

// ContainerName <- hash of current directory
func ContainerName() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return hash([]byte(currentDirectory))
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

func reset(context *cli.Context) {
	destroy(context)
	up(context)
	attach(context)
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
	docker("run", "-dti", volumesString(volumes), "--name", ContainerName(), imageID)
}

//
// CONFIG
//

func loadConfig() Config {
	configFile, err := os.Open(".workspace.json")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return config
}

func volumesString(volumes map[string]string) string {
	output := ""

	for host, dest := range volumes {
		absHostPath, err := filepath.Abs(host)
		if err != nil {
			panic(err)
		}
		volumeCommand := []string{output, "-v", absHostPath, ":", dest, ""}
		output = strings.Join(volumeCommand, "")
	}

	return output
}

//
// HELPER
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
