// Package actions provides functions that modify the surrounding context.
// functions in actions should not read information from the surrounding context
package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Attach spawns a bash shell in the container with given name
func Attach(containerName string) {
	cmd := exec.Command("docker", "exec", "-i", "-t", containerName, "/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// BuildImage builds Dockerfile specified in the config and returns the resulting Image ID
func BuildImage(dockerFileDirectory string) string {
	stdoutStderr := docker("build", "--quiet", dockerFileDirectory)
	imageID := strings.TrimSpace(strings.Split(string(stdoutStderr), ":")[1])

	fmt.Println("built image with ID:", imageID)
	fmt.Println("from Dockerfile in:", dockerFileDirectory)

	return imageID
}

// Destroy the container with the given name
func Destroy(containerName string) {
	docker("rm", "--force", containerName)
	fmt.Println("Destroyed container with name:", containerName)
}

// StartBackgroundContainer brings up a container with the given imageID and volume mappings
func StartBackgroundContainer(imageID string, name string, volumes map[string]string) {
	docker("run", "-dti", volumeArgs(volumes), "--name", name, imageID)
	fmt.Println("started background container withname:", name)
}

