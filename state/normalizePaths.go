package state

import (
	"path/filepath"
)

func makeHostPathsAbsolute(config TConfig) TConfig {
	config.DockerfileDirectory = makePathAbsolute(config.DockerfileDirectory)
	config.Volumes = makeKeysAbsolute(config.Volumes)
	return config
}

func makeKeysAbsolute(volumes map[string]string) map[string]string {
	normalizedVolumes := make(map[string]string)
	for host, guest := range volumes {
		normalizedVolumes[makePathAbsolute(host)] = guest
	}
	return normalizedVolumes
}

func makePathAbsolute(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(projectRoot(), path)
}
