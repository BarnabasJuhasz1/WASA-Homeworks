package util

import (
	"crypto/rand"
	"math/big"
)

const secureCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Mock database that maps tokens to user IDs (Username)
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
