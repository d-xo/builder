package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// BuildImage builds Dockerfile specified in the config and returns the resulting Image ID
func BuildImage(dockerFileDirectory string) string {
	stdoutStderr, err := exec.Command("docker", "build", "--quiet", dockerFileDirectory).CombinedOutput()
	if err != nil {
		fmt.Println(string(stdoutStderr))
		fmt.Println(err.Error())
		os.Exit(1)
	}

	imageID := strings.TrimSpace(strings.Split(string(stdoutStderr), ":")[1])

	fmt.Println("built image with ID:", imageID)
	fmt.Println("from Dockerfile in:", dockerFileDirectory)

	return imageID
}

// StartBackgroundContainer brings up a container with the given imageID and volume mappings
func StartBackgroundContainer(imageID string, name string, volumes map[string]string) {
	docker("run", "-dti", volumeArgs(volumes), "--name", name, imageID)
}

// Attach spawns a bash shell in the container with given name
func Attach(containerName string) {
	docker("exec", "-i", "-t", containerName, "/bin/bash")
}

// Destroy destroys the container with the given name
func Destroy(containerName string) {
	docker("rm", "--force", containerName)
	fmt.Println("Destroyed container with name:", containerName)
}

func docker(args ...string) {
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdoutStderr))
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func volumeArgs(volumes map[string]string) string {
	output := ""

	for host, dest := range volumes {
		volumeCommand := []string{output, "-v", host, ":", dest, " "}
		output = strings.Join(volumeCommand, "")
	}

	return strings.TrimSpace(output)
}
