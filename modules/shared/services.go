package shared

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("Failed to hash password %w", err)
	}

	return string(hashedPassword), nil
}

func (s *service) CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
