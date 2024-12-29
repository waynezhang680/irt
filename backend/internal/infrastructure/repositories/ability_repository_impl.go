package repositories

import (
	"context"
	"errors"

	"irt-exam-system/backend/internal/domain/repositories"
	"irt-exam-system/backend/models"
	"gorm.io/gorm"
)

type abilityRepository struct {
	db *gorm.DB
}

// NewAbilityRepository 创建能力值仓储实例
func NewAbilityRepository(db *gorm.DB) repositories.AbilityRepository {
	return &abilityRepository{db: db}
}

// 用户能力值操作实现
func (r *abilityRepository) CreateUserAbility(ctx context.Context, ability *models.UserAbility) error {
	return r.db.WithContext(ctx).Create(ability).Error
}

func (r *abilityRepository) UpdateUserAbility(ctx context.Context, ability *models.UserAbility) error {
	return r.db.WithContext(ctx).Save(ability).Error
}

func (r *abilityRepository) FindUserAbility(ctx context.Context, userID, subjectID uint) (*models.UserAbility, error) {
	var ability models.UserAbility
	err := r.db.WithContext(ctx).Where("user_id = ? AND subject_id = ?", userID, subjectID).First(&ability).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ability, nil
}

func (r *abilityRepository) ListUserAbilities(ctx context.Context, userID uint) ([]*models.UserAbility, error) {
	var abilities []*models.UserAbility
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&abilities).Error
	return abilities, err
}

func (r *abilityRepository) ListSubjectAbilities(ctx context.Context, subjectID uint, offset, limit int) ([]*models.UserAbility, int64, error) {
	var abilities []*models.UserAbility
	var total int64

	err := r.db.WithContext(ctx).Model(&models.UserAbility{}).Where("subject_id = ?", subjectID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("subject_id = ?", subjectID).
		Offset(offset).Limit(limit).Find(&abilities).Error
	if err != nil {
		return nil, 0, err
	}

	return abilities, total, nil
}

// 能力值估计记录操作实现
func (r *abilityRepository) CreateEstimation(ctx context.Context, estimation *models.AbilityEstimation) error {
	return r.db.WithContext(ctx).Create(estimation).Error
}

func (r *abilityRepository) ListUserEstimations(ctx context.Context, userID uint, offset, limit int) ([]*models.AbilityEstimation, int64, error) {
	var estimations []*models.AbilityEstimation
	var total int64

	err := r.db.WithContext(ctx).Model(&models.AbilityEstimation{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("user_id = ?", userID).
		Order("created_at DESC").Offset(offset).Limit(limit).Find(&estimations).Error
	if err != nil {
		return nil, 0, err
	}

	return estimations, total, nil
}

func (r *abilityRepository) ListSubjectEstimations(ctx context.Context, subjectID uint, offset, limit int) ([]*models.AbilityEstimation, int64, error) {
	var estimations []*models.AbilityEstimation
	var total int64

	err := r.db.WithContext(ctx).Model(&models.AbilityEstimation{}).Where("subject_id = ?", subjectID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("subject_id = ?", subjectID).
		Order("created_at DESC").Offset(offset).Limit(limit).Find(&estimations).Error
	if err != nil {
		return nil, 0, err
	}

	return estimations, total, nil
}

func (r *abilityRepository) GetLatestEstimation(ctx context.Context, userID, subjectID uint) (*models.AbilityEstimation, error) {
	var estimation models.AbilityEstimation
	err := r.db.WithContext(ctx).Where("user_id = ? AND subject_id = ?", userID, subjectID).
		Order("created_at DESC").First(&estimation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &estimation, nil
}

// 题目参数操作实现
func (r *abilityRepository) UpdateQuestionParameters(ctx context.Context, params *models.QuestionParameter) error {
	return r.db.WithContext(ctx).Save(params).Error
}

func (r *abilityRepository) GetQuestionParameters(ctx context.Context, questionID uint) (*models.QuestionParameter, error) {
	var params models.QuestionParameter
	err := r.db.WithContext(ctx).Where("question_id = ?", questionID).First(&params).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &params, nil
}

func (r *abilityRepository) BatchUpdateParameters(ctx context.Context, params []*models.QuestionParameter) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, param := range params {
			if err := tx.Save(param).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *abilityRepository) ListQuestionParameters(ctx context.Context, questionIDs []uint) ([]*models.QuestionParameter, error) {
	var params []*models.QuestionParameter
	err := r.db.WithContext(ctx).Where("question_id IN ?", questionIDs).Find(&params).Error
	return params, err
}
