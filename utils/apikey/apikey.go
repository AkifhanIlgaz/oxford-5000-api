package apikey

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateAPIKey creates a new secure random API key
// Returns a base64 encoded string in the format: "oxf_<random-string>"
func GenerateAPIKey() (string, error) {
	// Generate 32 bytes of random data
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("generate random bytes: %w", err)
	}

	// Encode to base64
	encoded := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Add prefix
	return fmt.Sprintf("oxf_%s", encoded), nil
}
