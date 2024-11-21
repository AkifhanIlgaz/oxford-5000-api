package crypto

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain text password and returns its bcrypt hash.
// It uses the default cost for the hashing algorithm.
// Returns the hashed password as a string and any error that occurred during hashing.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}

	return string(hashedPassword), nil
}

// VerifyPassword compares a hashed password with a plain text password.
// Returns nil if the passwords match, or an error if they don't.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
