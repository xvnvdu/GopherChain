package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Hash [32]byte

// Calculate hash
func NewHash(val any) Hash {
	jval, err := json.Marshal(val)
	if err != nil {
		panic("Failed to marshal for hashing:" + err.Error())
	}

	return sha256.Sum256(jval)
}

// Represent hash as a string
func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}
