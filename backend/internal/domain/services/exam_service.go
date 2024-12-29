package services

import (
	"context"
	"irt-exam-system/backend/internal/domain/models"
)

type ExamService interface {
	StartExam(ctx context.Context, userID uint) (*models.ExamSessionResponse, error)
	GetNextQuestion(ctx context.Context, sessionID uint) (*models.QuestionDTO, error)
	SubmitAnswer(ctx context.Context, sessionID uint, answer *models.AnswerRequest) (*models.AnswerResponse, error)
}
