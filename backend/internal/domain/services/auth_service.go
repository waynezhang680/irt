package services

import (
	"context"
	"irt-exam-system/backend/internal/domain/models"
)

type AuthService interface {
	Register(ctx context.Context, user *models.UserRegisterRequest) (*models.UserResponse, error)
	Login(ctx context.Context, credentials *models.LoginRequest) (*models.LoginResponse, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
}
