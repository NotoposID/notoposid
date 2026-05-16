package users

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) bool
}

type userService struct{}

func NewService() Service {
	return &userService{}
}

func (s *userService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *userService) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
