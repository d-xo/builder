package actions

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
	"os/exec"
	"strings"
)

func dockerClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

func IsContainerPresent(candidateName string) bool {
	allContainers, err := dockerClient().ContainerList(
		context.Background(),
		types.ContainerListOptions{
			All: true,
		},
	)
	if err != nil {
		panic(err)
	}

	return containerWithNameExists(candidateName, allContainers)
}

func containerWithNameExists(containerName string, containers []types.Container) bool {
	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return true
			}
		}
	}
	return false
}

func dockerCommandLine(args ...string) string {
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdoutStderr))
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(stdoutStderr)
}

func volumeArgs(volumes map[string]string) string {
	output := ""

	for host, dest := range volumes {
		volumeCommand := []string{output, "-v", host, ":", dest, " "}
		output = strings.Join(volumeCommand, "")
	}

	return strings.TrimSpace(output)
}
