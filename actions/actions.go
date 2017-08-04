// Package actions provides functions that modify the surrounding transforms.
// All state in actions should be passed in as a parameter (i.e. don't read info from the surroundings)
package actions

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

// Attach spawns a bash shell in the container with the given name
func Attach(containerName string) {
	cmd := exec.Command("docker", "exec", "-i", "-t", containerName, "/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// BuildImage builds Dockerfile specified in the config and returns the resulting Image ID
// run build twice, once to build the image (with stdout), and once to get the image ID
// this is UGLY
func BuildImage(dockerFileDirectory string) string {
	cmd := exec.Command("docker", "build", dockerFileDirectory)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	stdoutStderr := dockerCommandLine("build", "--quiet", dockerFileDirectory)
	imageID := strings.TrimSpace(strings.Split(string(stdoutStderr), ":")[1])

	fmt.Println("built image with ID:", imageID)
	fmt.Println("from Dockerfile in:", dockerFileDirectory)

	return imageID
}

// Destroy the container with the given name
func Destroy(containerName string) {
	dockerCommandLine("rm", "--force", containerName)
	fmt.Println("Destroyed container with name:", containerName)
}

// ExecuteDockerCommand runs a single docker command in the project build environment
func ExecuteDockerCommand(containerName string, command ...string) {
	args := append([]string{"exec", "-t", containerName}, command...)
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// StartBackgroundContainer brings up a container with the given imageID and volume mappings
func StartBackgroundContainer(imageID string, name string, volumes map[string]string) {
	dockerCommandLine("run", "-dti", volumeArgs(volumes), "--name", name, imageID)
	fmt.Println("started background container with name:", name)
}

// IsContainerPresent checks if a container with the given name is present on the system in any state (running, stopped etc...)
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

// IsContainerRunning checks if a container with the given name is running
func IsContainerRunning(candidateName string) bool {
	runningContainerFilters := filters.NewArgs()
	runningContainerFilters.Add("status", "running")

	runningContainers, err := dockerClient().ContainerList(
		context.Background(),
		types.ContainerListOptions{
			All:     true,
			Filters: runningContainerFilters,
		},
	)
	if err != nil {
		panic(err)
	}

	return containerWithNameExists(candidateName, runningContainers)
}
