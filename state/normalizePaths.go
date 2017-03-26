package state

import (
	"path/filepath"
)

// NormalizePath makes a path in .workspace.json absolute no matter the current working directory
func NormalizePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(projectRoot(), path)
}

func normalizeVolumes(volumes map[string]string) map[string]string {
	normalizedVolumes := make(map[string]string)
	for host, guest := range volumes {
		normalizedVolumes[NormalizePath(host)] = guest
	}
	return normalizedVolumes
}

func normalizeConfig(config TConfig) TConfig {
	config.DockerfileDirectory = NormalizePath(config.DockerfileDirectory)
	config.Volumes = normalizeVolumes(config.Volumes)
	return config
}
