package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

// build twice, once to show stdout, and once to get imageID
func buildImage(dockerFileDirectory string) string {
	cmd := exec.Command("docker", "image", "build", dockerFileDirectory)
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

func executeInContainer(containerName string, command ...string) {
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

func startBackgroundContainer(imageID string, name string, volumes map[string]string, privileged bool) {

	container := createContainer(imageID, name, volumes, privileged)

	client, ctx := dockerClient()
	if err := client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("Started background container with name:", name)
}

func isContainerPresent(candidateName string) bool {
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

func isContainerRunning(candidateName string) bool {
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
