package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/xwvvvvwx/workspace/state"
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

func StartBackgroundContainer(imageID string, volumes map[string]string) {
	Docker("run", "-dti", dockerArgsVolumesString(volumes), "--name", state.ContainerName(), imageID)
}

func Docker(args ...string) {
	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func dockerArgsVolumesString(volumes map[string]string) string {
	output := ""

	for host, dest := range volumes {
		volumeCommand := []string{output, "-v", state.NormalizePath(host), ":", dest, ""}
		output = strings.Join(volumeCommand, "")
	}

	return output
}
