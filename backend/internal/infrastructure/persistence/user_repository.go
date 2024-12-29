package persistence

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.DB.Create(user).Error
} 