// Package actions provides functions that modify the surrounding transforms.
// All state in actions should be passed in as a parameter (i.e. don't read info from the surroundings)
package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

// Attach spawns a shell in the container with the given name
func Attach(containerName string) {
	cmd := exec.Command("docker", "exec", "-it", containerName, "/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// BuildImage builds Dockerfile specified in the config and returns the resulting Image ID
// run build twice, once to build the image (with stdout), and once to get the image ID
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
	client, ctx := dockerClient()

	options := types.ContainerRemoveOptions{
		Force: true,
	}

	if err := client.ContainerRemove(ctx, containerName, options); err != nil {
		panic(err)
	}

	fmt.Println("Destroyed container with name:", containerName)
}

// ExecuteDockerCommand runs a single docker command in the project build environment
func ExecuteDockerCommand(containerName string, command ...string) {
	args := append([]string{"exec", "-t", containerName, "sh", "-c"}, command...)
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func createContainer(imageID string, name string, volumes map[string]string, privileged bool) container.ContainerCreateCreatedBody {
	client, ctx := dockerClient()

	var binds []string
	for host, dest := range volumes {
		binds = append(binds, host+":"+dest)
	}

	containerConfig := &container.Config{
		Image:     imageID,
		OpenStdin: true,
		Tty:       true,
	}
	hostConfig := &container.HostConfig{
		Binds:      binds,
		Privileged: privileged,
	}

	resp, err := client.ContainerCreate(ctx, containerConfig, hostConfig, nil, name)
	if err != nil {
		panic(err)
	}

	return resp
}

// StartBackgroundContainer brings up a container with the given imageID and volume mappings
func StartBackgroundContainer(imageID string, name string, volumes map[string]string, privileged bool) {

	resp := createContainer(imageID, name, volumes, privileged)

	client, ctx := dockerClient()
	if err := client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("Started background container with name:", name)
}

// IsContainerPresent checks if a container with the given name is present on the system in any state (running, stopped etc...)
func IsContainerPresent(candidateName string) bool {
	client, ctx := dockerClient()

	allContainers, err := client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return containerWithNameExists(candidateName, allContainers)
}

func running() filters.Args {
	filter := filters.NewArgs()
	filter.Add("status", "running")
	return filter
}

// IsContainerRunning checks if a container with the given name is running
func IsContainerRunning(candidateName string) bool {

	client, ctx := dockerClient()

	runningContainers, err := client.ContainerList(
		ctx,
		types.ContainerListOptions{All: true, Filters: running()},
	)
	if err != nil {
		panic(err)
	}

	return containerWithNameExists(candidateName, runningContainers)
}
