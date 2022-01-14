package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash gets a string and returns its hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword checks if the password matches with a given hash
func VerifyPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
