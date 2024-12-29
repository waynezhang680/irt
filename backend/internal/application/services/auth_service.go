package services

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repositories.UserRepository
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
} 