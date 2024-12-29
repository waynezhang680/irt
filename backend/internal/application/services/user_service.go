package services

import (
	"context"

	"irt-exam-system/backend/internal/domain/repositories"
	"irt-exam-system/backend/models"
)

// UserService 用户服务接口
type UserService interface {
	// 用户管理
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	GetUser(ctx context.Context, id uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	ListUsers(ctx context.Context, offset, limit int) ([]*models.User, int64, error)

	// 认证相关
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, token string) error
	ValidateToken(ctx context.Context, token string) (*models.User, error)
	RefreshToken(ctx context.Context, token string) (string, error)
	ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// 实现UserService接口
func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *userService) GetUser(ctx context.Context, id uint) (*models.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.FindByEmail(ctx, email)
}

func (s *userService) ListUsers(ctx context.Context, offset, limit int) ([]*models.User, int64, error) {
	return s.userRepo.List(ctx, offset, limit)
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	// TODO: 实现登录逻辑，包括密码验证和JWT token生成
	return "", nil
}

func (s *userService) Logout(ctx context.Context, token string) error {
	// TODO: 实现登出逻辑，包括token黑名单
	return nil
}

func (s *userService) ValidateToken(ctx context.Context, token string) (*models.User, error) {
	// TODO: 实现token验证逻辑
	return nil, nil
}

func (s *userService) RefreshToken(ctx context.Context, token string) (string, error) {
	// TODO: 实现token刷新逻辑
	return "", nil
}

func (s *userService) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	// TODO: 实现密码修改逻辑
	return nil
}
