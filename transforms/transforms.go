// Package transforms provides functions gather information from the surrounding context
// functions in transforms should be pure
package transforms

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

// TConfig enforces the .workspace.json schema
type TConfig struct {
	DockerfileDirectory string            `json:"dockerfileDirectory"`
	Volumes             map[string]string `json:"volumes"`
	Alias               map[string]string `json:"commands"`
}

// Config reads the .workspace.json
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
	hasher := sha1.New()
	hasher.Write([]byte(projectRoot()))
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}

// CommandFromAlias searches for the given alias in the .workspace.json
func CommandFromAlias(aliasName string) string {
	value, present := Config().Alias[aliasName]
	if !present {
		fmt.Println(aliasName, "was not found in", configPath())
		os.Exit(1)
	}
	fmt.Println(value)
	return value
}
