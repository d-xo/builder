package state

import (
	"crypto/sha1"
	"encoding/hex"
)

func hash(bytes []byte) string {
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}
