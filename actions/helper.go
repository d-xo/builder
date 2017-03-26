package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func docker(args ...string) string {
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
