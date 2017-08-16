package commands

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/docker/docker/client"
)

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

func dockerClient() (*client.Client, context.Context) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli, ctx
}
