package repositories

import (
	"context"

	"irt-exam-system/backend/models"
)

// AbilityRepository 能力值仓储接口
type AbilityRepository interface {
	// 用户能力值操作
	CreateUserAbility(ctx context.Context, ability *models.UserAbility) error
	UpdateUserAbility(ctx context.Context, ability *models.UserAbility) error
	FindUserAbility(ctx context.Context, userID, subjectID uint) (*models.UserAbility, error)
	ListUserAbilities(ctx context.Context, userID uint) ([]*models.UserAbility, error)
	ListSubjectAbilities(ctx context.Context, subjectID uint, offset, limit int) ([]*models.UserAbility, int64, error)

	// 能力值估计记录操作
	CreateEstimation(ctx context.Context, estimation *models.AbilityEstimation) error
	ListUserEstimations(ctx context.Context, userID uint, offset, limit int) ([]*models.AbilityEstimation, int64, error)
	ListSubjectEstimations(ctx context.Context, subjectID uint, offset, limit int) ([]*models.AbilityEstimation, int64, error)
	GetLatestEstimation(ctx context.Context, userID, subjectID uint) (*models.AbilityEstimation, error)

	// 题目参数操作
	UpdateQuestionParameters(ctx context.Context, params *models.QuestionParameter) error
	GetQuestionParameters(ctx context.Context, questionID uint) (*models.QuestionParameter, error)
	BatchUpdateParameters(ctx context.Context, params []*models.QuestionParameter) error
	ListQuestionParameters(ctx context.Context, questionIDs []uint) ([]*models.QuestionParameter, error)
}
