package services

import (
	"context"

	"irt-exam-system/backend/internal/domain/repositories"
	"irt-exam-system/backend/models"
)

// KnowledgeService 知识点服务接口
type KnowledgeService interface {
	// 知识点管理
	CreateKnowledgePoint(ctx context.Context, point *models.KnowledgePoint) error
	UpdateKnowledgePoint(ctx context.Context, point *models.KnowledgePoint) error
	DeleteKnowledgePoint(ctx context.Context, id uint) error
	GetKnowledgePoint(ctx context.Context, id uint) (*models.KnowledgePoint, error)
	ListKnowledgePoints(ctx context.Context, offset, limit int) ([]*models.KnowledgePoint, int64, error)
	GetSubjectKnowledgePoints(ctx context.Context, subjectID uint) ([]*models.KnowledgePoint, error)

	// 知识点关联
	AddQuestionToKnowledgePoint(ctx context.Context, pointID, questionID uint) error
	RemoveQuestionFromKnowledgePoint(ctx context.Context, pointID, questionID uint) error
	GetKnowledgePointQuestions(ctx context.Context, pointID uint) ([]*models.Question, error)
}

// KnowledgePointProgress 知识点掌握进度
type KnowledgePointProgress struct {
	KnowledgePointID  uint    `json:"knowledge_point_id"`
	Name              string  `json:"name"`
	MasteryLevel      float64 `json:"mastery_level"`
	QuestionCount     int     `json:"question_count"`
	CorrectCount      int     `json:"correct_count"`
	AverageTimeSpent  float64 `json:"average_time_spent"`
	RecommendedReview bool    `json:"recommended_review"`
}

// NewKnowledgeService creates a new knowledge service instance
func NewKnowledgeService(knowledgeRepo repositories.KnowledgePointRepository) KnowledgeService {
	return &knowledgeService{
		knowledgeRepo: knowledgeRepo,
	}
}

type knowledgeService struct {
	knowledgeRepo repositories.KnowledgePointRepository
}

// CreateKnowledgePoint implements KnowledgeService
func (s *knowledgeService) CreateKnowledgePoint(ctx context.Context, point *models.KnowledgePoint) error {
	return s.knowledgeRepo.Create(ctx, point)
}

// UpdateKnowledgePoint implements KnowledgeService
func (s *knowledgeService) UpdateKnowledgePoint(ctx context.Context, point *models.KnowledgePoint) error {
	return s.knowledgeRepo.Update(ctx, point)
}

// DeleteKnowledgePoint implements KnowledgeService
func (s *knowledgeService) DeleteKnowledgePoint(ctx context.Context, id uint) error {
	return s.knowledgeRepo.Delete(ctx, id)
}

// GetKnowledgePoint implements KnowledgeService
func (s *knowledgeService) GetKnowledgePoint(ctx context.Context, id uint) (*models.KnowledgePoint, error) {
	return s.knowledgeRepo.FindByID(ctx, id)
}

// ListKnowledgePoints implements KnowledgeService
func (s *knowledgeService) ListKnowledgePoints(ctx context.Context, offset, limit int) ([]*models.KnowledgePoint, int64, error) {
	points, err := s.knowledgeRepo.ListByParent(ctx, 0) // 获取顶级知识点
	if err != nil {
		return nil, 0, err
	}
	return points, int64(len(points)), nil
}

// GetSubjectKnowledgePoints implements KnowledgeService
func (s *knowledgeService) GetSubjectKnowledgePoints(ctx context.Context, subjectID uint) ([]*models.KnowledgePoint, error) {
	return s.knowledgeRepo.ListBySubject(ctx, subjectID)
}

// AddQuestionToKnowledgePoint implements KnowledgeService
func (s *knowledgeService) AddQuestionToKnowledgePoint(ctx context.Context, pointID, questionID uint) error {
	return s.knowledgeRepo.AddQuestion(ctx, pointID, questionID)
}

// RemoveQuestionFromKnowledgePoint implements KnowledgeService
func (s *knowledgeService) RemoveQuestionFromKnowledgePoint(ctx context.Context, pointID, questionID uint) error {
	return s.knowledgeRepo.RemoveQuestion(ctx, pointID, questionID)
}

// GetKnowledgePointQuestions implements KnowledgeService
func (s *knowledgeService) GetKnowledgePointQuestions(ctx context.Context, pointID uint) ([]*models.Question, error) {
	questions, _, err := s.knowledgeRepo.ListQuestions(ctx, pointID, 0, -1)
	return questions, err
}
