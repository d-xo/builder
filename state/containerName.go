package state

import (
	"crypto/sha1"
	"encoding/hex"
)

// ContainerName is the hash of the project root path
func ContainerName() string {
	return hash([]byte(projectRoot()))
}

func hash(bytes []byte) string {
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}
