package repositories

import (
	"context"

	"irt-exam-system/backend/internal/domain/models"
)

// ExamRepository 考试仓储接口
type ExamRepository interface {
	// 试卷相关
	CreatePaper(ctx context.Context, paper *models.ExamPaper) error
	UpdatePaper(ctx context.Context, paper *models.ExamPaper) error
	DeletePaper(ctx context.Context, id uint) error
	FindPaperByID(ctx context.Context, id uint) (*models.ExamPaper, error)
	ListPapers(ctx context.Context, offset, limit int) ([]*models.ExamPaper, int64, error)
	ListPapersBySubject(ctx context.Context, subjectID uint, offset, limit int) ([]*models.ExamPaper, int64, error)
	ListPapersByStatus(ctx context.Context, status string, offset, limit int) ([]*models.ExamPaper, int64, error)

	// 考试记录相关
	CreateRecord(ctx context.Context, record *models.ExamSession) error
	UpdateRecord(ctx context.Context, record *models.ExamSession) error
	FindRecordByID(ctx context.Context, id uint) (*models.ExamSession, error)
	ListUserRecords(ctx context.Context, userID uint, offset, limit int) ([]*models.ExamSession, int64, error)
	ListPaperRecords(ctx context.Context, paperID uint, offset, limit int) ([]*models.ExamSession, int64, error)

	// 答题记录相关
	CreateResponse(ctx context.Context, response *models.QuestionResponse) error
	BatchCreateResponses(ctx context.Context, responses []*models.QuestionResponse) error
	UpdateResponse(ctx context.Context, response *models.QuestionResponse) error
	ListRecordResponses(ctx context.Context, recordID uint) ([]*models.QuestionResponse, error)
	GetResponseStats(ctx context.Context, questionID uint) (total int64, correct int64, error error)
}
