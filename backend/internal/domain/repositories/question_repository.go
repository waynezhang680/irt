package repositories

import (
	"context"

	"irt-exam-system/backend/internal/domain/models"
)

// QuestionRepository 试题仓储接口
type QuestionRepository interface {
	// 基本操作
	Create(ctx context.Context, question *models.Question) error
	Update(ctx context.Context, question *models.Question) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*models.Question, error)

	// 查询操作
	ListBySubject(ctx context.Context, subjectID uint, offset, limit int) ([]*models.Question, int64, error)
	ListByKnowledgePoint(ctx context.Context, knowledgePointID uint, offset, limit int) ([]*models.Question, int64, error)
	ListByType(ctx context.Context, questionType string, offset, limit int) ([]*models.Question, int64, error)
	Search(ctx context.Context, keyword string, offset, limit int) ([]*models.Question, int64, error)
	ListByExamPaper(ctx context.Context, examPaperID uint) ([]*models.Question, error)
	FindByDifficulty(ctx context.Context, difficulty float64) (*models.Question, error)

	// 选项操作
	CreateOption(ctx context.Context, option *models.QuestionOption) error
	UpdateOption(ctx context.Context, option *models.QuestionOption) error
	DeleteOption(ctx context.Context, id uint) error
	ListOptions(ctx context.Context, questionID uint) ([]*models.QuestionOption, error)

	// 知识点关联
	AddKnowledgePoint(ctx context.Context, questionID, knowledgePointID uint) error
	RemoveKnowledgePoint(ctx context.Context, questionID, knowledgePointID uint) error
	ListKnowledgePoints(ctx context.Context, questionID uint) ([]*models.KnowledgePoint, error)

	// IRT参数操作
	UpdateParameters(ctx context.Context, params *models.QuestionParameter) error
	GetParameters(ctx context.Context, questionID uint) (*models.QuestionParameter, error)
	BatchGetParameters(ctx context.Context, questionIDs []uint) ([]*models.QuestionParameter, error)
}
