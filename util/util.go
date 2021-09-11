package util

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHexString(lenInBytes int) (string, error) {
	buf := make([]byte, lenInBytes)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(buf), nil
}
