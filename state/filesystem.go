package state

import (
	"fmt"
	"os"
	"path/filepath"
	"crypto/sha1"
	"encoding/hex"
)

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

func hash(bytes []byte) string {
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}

func configPath() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
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
