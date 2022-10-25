package domain

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Password string
type PasswordHash string

func (p Password) Hash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	if err != nil {
		return "", fmt.Errorf("password hashing error: %w", err)
	}
	return string(bytes), err
}

func (p Password) CheckPasswordHash(hash PasswordHash) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
