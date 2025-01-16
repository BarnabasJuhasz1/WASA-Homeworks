package util

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const secureCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// maps tokens to user IDs (Username)
var TokenMap = map[string]string{}

func GetRandomToken(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		// Generate a secure random index in the charset
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(secureCharset))))
		if err != nil {
			return "", err
		}
		b[i] = secureCharset[idx.Int64()]
	}
	return string(b), nil
}

// UpdateUsername updates the username associated with a token in the map
func UpdateUsername(oldUsername string, newUsername string) error {
	// Find the token associated with the old username
	var foundToken string
	for token, username := range TokenMap {
		if username == oldUsername {
			foundToken = token
			break
		}
	}

	// If no matching token is found, return an error
	if foundToken == "" {
		return errors.New("username not found in token map")
	}

	// msg := "new assoc. for token: " + foundToken + " is now: " + newUsername + "  (old: " + oldUsername + ")"

	// Update the map with the new username while keeping the token unchanged
	TokenMap[foundToken] = newUsername

	return nil // errors.New(msg)
}
