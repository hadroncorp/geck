package hashing

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
)

// NewHashString encodes and hashes the given value.
// Routine is concurrent-safe.
func NewHashString(val any) (string, error) {
	encodingBuffer := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(encodingBuffer).Encode(val); err != nil {
		return "", errors.New("failed to encode value")
	}

	hasher := sha256.New()
	if _, err := io.Copy(hasher, encodingBuffer); err != nil {
		return "", errors.New("failed to hash value")
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
