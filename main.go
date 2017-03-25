package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

// -------------------------------------------------------------------------------------------------

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

// -------------------------------------------------------------------------------------------------
//  commands
// -------------------------------------------------------------------------------------------------

func up(context *cli.Context) {
	imageID := buildImage()
	startBackgroundContainer(imageID)
}

func attach(context *cli.Context) {
	docker("exec", "-i", "-t", containerName(), "/bin/bash")
}

func destroy(context *cli.Context) {
	docker("rm", "--force", containerName())
}

func reset(context *cli.Context) {
	destroy(context)
	up(context)
	attach(context)
}

// -------------------------------------------------------------------------------------------------
//  data
// -------------------------------------------------------------------------------------------------

func containerName() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Println("ERROR: Could not get current working directory")
		log.Fatal(err)
	}

	return hash([]byte(currentDirectory))
}

func hash(bytes []byte) string {
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}

// -------------------------------------------------------------------------------------------------
//  detail
// -------------------------------------------------------------------------------------------------

func buildImage() string {
	stdoutStderr, err := exec.Command("docker", "build", "-q", ".").CombinedOutput()
	if err != nil {
		log.Println(string(stdoutStderr))
		log.Fatal(err)
	}
	imageID := strings.Split(string(stdoutStderr), ":")[1]
	fmt.Println("built image with ID:", imageID)
	return strings.TrimSpace(imageID)
}

func startBackgroundContainer(imageID string) {
	docker("run", "-dti", "--name", containerName(), imageID)
}

func docker(args ...string) {
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
