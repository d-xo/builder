// Package data provides functions that gather information from the surroundings
// functions in data should not modify the surroundings in any way
package data

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

// TConfig enforces the .builder.json schema
type TConfig struct {
	DockerfileDirectory string            `json:"dockerfileDirectory"`
	Volumes             map[string]string `json:"volumes"`
	Alias               map[string]string `json:"commands"`
	ContainerName       string            `json:"containerName"`
}

// Config reads the .builder.json
func Config() TConfig {
	configFile, err := os.Open(configPath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var config TConfig
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return makeHostPathsAbsolute(config)
}

// ContainerName is hashed from the project root
func ContainerName() string {
	if Config().ContainerName != "" {
		return Config().ContainerName
	}

	hasher := sha1.New()
	hasher.Write([]byte(projectRoot()))
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}

// CommandFromAlias searches for the given alias in the .builder.json
func CommandFromAlias(aliasName string) string {
	value, present := Config().Alias[aliasName]
	if !present {
		fmt.Println(aliasName, "was not found in", configPath())
		os.Exit(1)
	}
	fmt.Println(value)
	return value
}
