/*
state computes abunch of
*/
package state

import (
	"encoding/json"
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
		panic(err)
	}

	var config TConfig
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panic(err)
	}

	return normalizeConfig(config)
}

// ContainerName is the hash of ProjectRoot()
func ContainerName() string {
	return hash([]byte(projectRoot()))
}
