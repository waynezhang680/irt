package services

import (
	"context"
	"math"

	"irt-exam-system/backend/internal/domain/repositories"
	"irt-exam-system/backend/models"
)

// AbilityService 能力值服务接口
type AbilityService interface {
	// 能力值管理
	GetUserAbility(ctx context.Context, userID, subjectID uint) (*models.UserAbility, error)
	ListUserAbilities(ctx context.Context, userID uint) ([]*models.UserAbility, error)
	ListSubjectAbilities(ctx context.Context, subjectID uint, offset, limit int) ([]*models.UserAbility, int64, error)

	// 能力值估计历史
	GetUserEstimationHistory(ctx context.Context, userID uint, offset, limit int) ([]*models.AbilityEstimation, int64, error)
	GetLatestEstimation(ctx context.Context, userID, subjectID uint) (*models.AbilityEstimation, error)

	// IRT参数管理
	UpdateQuestionParameters(ctx context.Context, questionID uint, difficulty, discrimination, guessing float64) error
	GetQuestionParameters(ctx context.Context, questionID uint) (*models.QuestionParameter, error)
	BatchUpdateParameters(ctx context.Context, params []*models.QuestionParameter) error

	// IRT模型计算
	CalculateResponseProbability(ctx context.Context, ability float64, questionID uint) (float64, error)
	EstimateAbility(ctx context.Context, responses []*models.ExamResponse) (float64, float64, error)
	GetConfidenceInterval(ctx context.Context, ability, standardError float64) (float64, float64, error)
	GetPerformanceLevel(ctx context.Context, ability float64) (string, error)
	GenerateRecommendations(ctx context.Context, userID uint, ability float64) ([]string, error)
}

// NewAbilityService creates a new ability service instance
func NewAbilityService(abilityRepo repositories.AbilityRepository) AbilityService {
	return &abilityService{
		abilityRepo: abilityRepo,
	}
}

type abilityService struct {
	abilityRepo repositories.AbilityRepository
}

// GetUserAbility implements AbilityService
func (s *abilityService) GetUserAbility(ctx context.Context, userID, subjectID uint) (*models.UserAbility, error) {
	return s.abilityRepo.FindUserAbility(ctx, userID, subjectID)
}

// ListUserAbilities implements AbilityService
func (s *abilityService) ListUserAbilities(ctx context.Context, userID uint) ([]*models.UserAbility, error) {
	return s.abilityRepo.ListUserAbilities(ctx, userID)
}

// ListSubjectAbilities implements AbilityService
func (s *abilityService) ListSubjectAbilities(ctx context.Context, subjectID uint, offset, limit int) ([]*models.UserAbility, int64, error) {
	return s.abilityRepo.ListSubjectAbilities(ctx, subjectID, offset, limit)
}

// GetUserEstimationHistory implements AbilityService
func (s *abilityService) GetUserEstimationHistory(ctx context.Context, userID uint, offset, limit int) ([]*models.AbilityEstimation, int64, error) {
	return s.abilityRepo.ListUserEstimations(ctx, userID, offset, limit)
}

// GetLatestEstimation implements AbilityService
func (s *abilityService) GetLatestEstimation(ctx context.Context, userID, subjectID uint) (*models.AbilityEstimation, error) {
	return s.abilityRepo.GetLatestEstimation(ctx, userID, subjectID)
}

// UpdateQuestionParameters implements AbilityService
func (s *abilityService) UpdateQuestionParameters(ctx context.Context, questionID uint, difficulty, discrimination, guessing float64) error {
	params := &models.QuestionParameter{
		QuestionID:     questionID,
		Difficulty:     difficulty,
		Discrimination: discrimination,
		Guessing:       guessing,
	}
	return s.abilityRepo.UpdateQuestionParameters(ctx, params)
}

// GetQuestionParameters implements AbilityService
func (s *abilityService) GetQuestionParameters(ctx context.Context, questionID uint) (*models.QuestionParameter, error) {
	return s.abilityRepo.GetQuestionParameters(ctx, questionID)
}

// BatchUpdateParameters implements AbilityService
func (s *abilityService) BatchUpdateParameters(ctx context.Context, params []*models.QuestionParameter) error {
	return s.abilityRepo.BatchUpdateParameters(ctx, params)
}

// CalculateResponseProbability implements AbilityService
func (s *abilityService) CalculateResponseProbability(ctx context.Context, ability float64, questionID uint) (float64, error) {
	params, err := s.abilityRepo.GetQuestionParameters(ctx, questionID)
	if err != nil {
		return 0, err
	}

	// 使用3PL模型计算作答概率
	// P(θ) = c + (1-c)/(1+e^(-1.7a(θ-b)))
	// θ: 能力值
	// a: 区分度
	// b: 难度
	// c: 猜测参数
	exp := -1.7 * params.Discrimination * (ability - params.Difficulty)
	prob := params.Guessing + (1-params.Guessing)/(1+math.Exp(exp))
	return prob, nil
}

// EstimateAbility implements AbilityService
func (s *abilityService) EstimateAbility(ctx context.Context, responses []*models.ExamResponse) (float64, float64, error) {
	// TODO: 实现最大似然估计(MLE)或贝叶斯估计(EAP)
	return 0, 0, nil
}

// GetConfidenceInterval implements AbilityService
func (s *abilityService) GetConfidenceInterval(ctx context.Context, ability, standardError float64) (float64, float64, error) {
	// 95%置信区间
	lower := ability - 1.96*standardError
	upper := ability + 1.96*standardError
	return lower, upper, nil
}

// GetPerformanceLevel implements AbilityService
func (s *abilityService) GetPerformanceLevel(ctx context.Context, ability float64) (string, error) {
	switch {
	case ability >= 2.0:
		return "优秀", nil
	case ability >= 1.0:
		return "良好", nil
	case ability >= -1.0:
		return "中等", nil
	case ability >= -2.0:
		return "及格", nil
	default:
		return "不及格", nil
	}
}

// GenerateRecommendations implements AbilityService
func (s *abilityService) GenerateRecommendations(ctx context.Context, userID uint, ability float64) ([]string, error) {
	// TODO: 基于能力值生成学习建议
	return []string{
		"继续保持",
		"多做练习",
		"查漏补缺",
	}, nil
}
