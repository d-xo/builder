package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/urfave/cli"
	"github.com/xwvvvvwx/builder/config"
)

// Up starts a background container from the Dockerfile specified in the config
func Up(c *cli.Context) {
	if isContainerPresent(config.ContainerName()) {
		Destroy(c)
	}

	imageID := buildImage(config.Config().DockerfileDirectory)
	startBackgroundContainer(
		imageID, config.ContainerName(),
		config.Config().Volumes,
		config.Config().Privileged,
	)
}

// Exec executes a single command in the build environment
func Exec(c *cli.Context) {
	if !isContainerRunning(config.ContainerName()) {
		Up(c)
	}
	command := append([]string{c.Args().First()}, c.Args().Tail()...)
	executeInContainer(config.ContainerName(), command...)
}

// Run executes the specified alias
func Run(c *cli.Context) {
	aliasName := c.Args().First()
	executeAlias(c, aliasName)
}

// Attach spawns a bash shell in the build environment
func Attach(c *cli.Context) {
	cmd := exec.Command("docker", "exec", "-it", config.ContainerName(), "/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// Destroy destroys the build environment
func Destroy(c *cli.Context) {
	client, ctx := dockerClient()

	options := types.ContainerRemoveOptions{
		Force: true,
	}

	if err := client.ContainerRemove(ctx, config.ContainerName(), options); err != nil {
		panic(err)
	}

	fmt.Println("Destroyed container with name:", config.ContainerName())
}

// Clean resets the build environment to the state specified in the Dockerfile
func Clean(c *cli.Context) {
	Destroy(c)
	Up(c)
}

// Build executes the "build" alias
func Build(c *cli.Context) {
	executeAlias(c, "build")
}

// Start executes the "start" alias
func Start(c *cli.Context) {
	executeAlias(c, "start")
}

// Verify executes the "verify" alias
func Verify(c *cli.Context) {
	executeAlias(c, "verify")
}

// Package executes the "package" alias
func Package(c *cli.Context) {
	executeAlias(c, "package")
}

// Benchmark executes the "benchmark" alias
func Benchmark(c *cli.Context) {
	executeAlias(c, "benchmark")
}

func executeAlias(c *cli.Context, aliasName string) {
	if !isContainerRunning(config.ContainerName()) {
		Up(c)
	}
	executeInContainer(config.ContainerName(), config.CommandFromAlias(aliasName))
}
