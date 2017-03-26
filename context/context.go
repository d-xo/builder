// Package context provides functions gather information from the surrounding context
// functions in context should not modify the context in any way
package context

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
