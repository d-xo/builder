package state

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// TConfig <- .workspace.json schema
type TConfig struct {
	DockerfileDirectory string            `json:"dockerfileDirectory"`
	Volumes             map[string]string `json:"volumes"`
}

// Config <- read .workspace.json from disk
func Config() TConfig {
	configFile, err := os.Open(ConfigPath())
	if err != nil {
		panic(err)
	}

	var config TConfig
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panic(err)
	}

	return normalizeConfig(config)
}

// ConfigPath <- recurse upwards until .workspace.json is found
func ConfigPath() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return findConfig(currentDirectory)
}

// ConfigDir <- directory of config file
func ConfigDir() string {
	return filepath.Dir(ConfigPath())
}

// ContainerName <- hash of ConfigPath()
func ContainerName() string {
	return hash([]byte(ConfigPath()))
}
