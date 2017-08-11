package actions

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func dockerClient() (*client.Client, context.Context) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli, ctx
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
