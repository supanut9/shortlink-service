package utils

import (
	"crypto/rand"
	"io"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateSlug creates a random, URL-safe slug of a given length.
// It uses crypto/rand for better security and uniqueness.
func GenerateSlug(length int) string {
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		// This should not happen in a standard environment
		panic("failed to read random bytes: " + err.Error())
	}

	for i := 0; i < length; i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}
