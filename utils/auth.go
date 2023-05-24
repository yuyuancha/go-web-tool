package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash generate a hash encrypted string
func GenerateHash(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CompareHash compare input and hash consistent
func CompareHash(input, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	if err != nil {
		return false, err
	}

	return true, nil
}
