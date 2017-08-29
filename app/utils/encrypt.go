package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword -
func EncryptPassword(s string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(s), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}
