package docker

import (
	"fmt"
	"github.com/xwvvvvwx/workspace/state"
	"os"
	"os/exec"
	"strings"
)

// BuildImage builds Dockerfile specified in the config and returns the resulting Image ID
func BuildImage(config state.TConfig) string {
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

// StartBackgroundContainer brings up a container with the given imageID and volume mappings
func StartBackgroundContainer(imageID string, volumes map[string]string) {
	Docker("run", "-dti", volumesString(volumes), "--name", state.ContainerName(), imageID)
}

// Docker executes docker with the given args
func Docker(args ...string) {
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func volumesString(volumes map[string]string) string {
	output := ""

	for host, dest := range volumes {
		volumeCommand := []string{output, "-v", state.NormalizePath(host), ":", dest, " "}
		output = strings.Join(volumeCommand, "")
	}

	return strings.TrimSpace(output)
}
