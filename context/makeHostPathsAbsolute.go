package context

import (
	"path/filepath"
)

func makeHostPathsAbsolute(config TConfig) TConfig {
	config.DockerfileDirectory = makePathAbsolute(config.DockerfileDirectory)
	config.Volumes = makeKeysAbsolutePaths(config.Volumes)
	return config
}

func makeKeysAbsolutePaths(volumes map[string]string) map[string]string {
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
