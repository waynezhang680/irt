package repositories

import (
	"context"
	"irt-exam-system/backend/internal/domain/models"
)

type ExamSessionRepository interface {
	Create(ctx context.Context, session *models.ExamSession) error
	FindByID(ctx context.Context, id uint) (*models.ExamSession, error)
	SaveResponse(ctx context.Context, response *models.QuestionResponse) error
	UpdateAbility(ctx context.Context, sessionID uint, newAbility float64) error
	GetResponses(ctx context.Context, sessionID uint) ([]*models.QuestionResponse, error)
}
