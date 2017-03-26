package context

import (
	"fmt"
	"os"
	"path/filepath"
)

func configPath() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return findConfig(currentDirectory)
}

func projectRoot() string {
	return filepath.Dir(configPath())
}

func findConfig(directory string) string {
	configFile := filepath.Join(directory, ".workspace.json")

	if pathExists(configFile) {
		return configFile
	}

	if isPathRoot(directory) {
		fmt.Println("No .workspace.json found")
		os.Exit(1)
	}

	return findConfig(filepath.Dir(directory))
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func isPathRoot(path string) bool {
	if filepath.Dir(path) == filepath.Dir(filepath.Dir(path)) {
		return true
	}
	return false
}
