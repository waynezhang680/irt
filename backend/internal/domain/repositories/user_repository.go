package repositories

import (
	"backend/internal/domain/models"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
}
