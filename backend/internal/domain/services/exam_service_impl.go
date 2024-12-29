package services

import (
	"context"
	"errors"
	"time"

	"irt-exam-system/backend/internal/domain/models"
	"irt-exam-system/backend/internal/domain/repositories"
	"irt-exam-system/backend/pkg/utils"
)

type ExamServiceImpl struct {
	questionRepo    repositories.QuestionRepository
	examSessionRepo repositories.ExamSessionRepository
	irtService      IRTService
}

func NewExamService(
	questionRepo repositories.QuestionRepository,
	examSessionRepo repositories.ExamSessionRepository,
	irtService IRTService,
) ExamService {
	return &ExamServiceImpl{
		questionRepo:    questionRepo,
		examSessionRepo: examSessionRepo,
		irtService:      irtService,
	}
}

func (s *ExamServiceImpl) StartExam(ctx context.Context, userID uint) (*models.ExamSessionResponse, error) {
	session := &models.ExamSession{
		UserID:         userID,
		StartTime:      time.Now(),
		CurrentAbility: 0.0, // 初始能力值设为0
		Status:         "in_progress",
	}

	if err := s.examSessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	return &models.ExamSessionResponse{
		ID:             session.ID,
		StartTime:      session.StartTime.Format(time.RFC3339),
		CurrentAbility: session.CurrentAbility,
	}, nil
}

func (s *ExamServiceImpl) GetNextQuestion(ctx context.Context, sessionID uint) (*models.QuestionDTO, error) {
	session, err := s.examSessionRepo.FindByID(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	if session.Status != "in_progress" {
		return nil, errors.New("exam session is not in progress")
	}

	// 根据当前能力值选择合适难度的题目
	question, err := s.questionRepo.FindByDifficulty(ctx, session.CurrentAbility)
	if err != nil {
		return nil, errors.New("no suitable questions found")
	}

	return &models.QuestionDTO{
		ID:      question.ID,
		Content: question.Content,
		Options: []string{
			question.OptionA,
			question.OptionB,
			question.OptionC,
			question.OptionD,
		},
	}, nil
}

func (s *ExamServiceImpl) SubmitAnswer(ctx context.Context, sessionID uint, answer *models.AnswerRequest) (*models.AnswerResponse, error) {
	session, err := s.examSessionRepo.FindByID(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	question, err := s.questionRepo.FindByID(ctx, answer.QuestionID)
	if err != nil {
		return nil, err
	}

	// 记录答题结果
	isCorrect := answer.Answer == question.Answer
	response := &models.QuestionResponse{
		ExamSessionID: sessionID,
		QuestionID:    question.ID,
		UserAnswer:    answer.Answer,
		IsCorrect:     isCorrect,
		TimeSpent:     answer.TimeSpent,
	}
	if err := s.examSessionRepo.SaveResponse(ctx, response); err != nil {
		return nil, err
	}

	// 更新考生能力值估计
	newAbility := s.irtService.EstimateAbility(
		session.CurrentAbility,
		question.Difficulty,
		question.Discrimination,
		question.GuessParameter,
		isCorrect,
	)

	if err := s.examSessionRepo.UpdateAbility(ctx, sessionID, newAbility); err != nil {
		return nil, err
	}

	// 计算下一题的建议难度
	nextDifficulty := s.irtService.GetNextQuestionDifficulty(newAbility)

	return &models.AnswerResponse{
		IsCorrect:      isCorrect,
		CurrentAbility: newAbility,
		NextDifficulty: nextDifficulty,
	}, nil
}

// calculateTestInformation 计算整个测验的信息量
func (s *ExamServiceImpl) calculateTestInformation(ctx context.Context, ability float64, questions []*models.Question) float64 {
	var totalInfo float64
	for _, q := range questions {
		info := utils.CalculateItemInformation(
			ability,
			q.Difficulty,
			q.Discrimination,
			q.GuessParameter,
		)
		totalInfo += info
	}
	return totalInfo
}

// shouldStopTest 判断是否应该结束测验
func (s *ExamServiceImpl) shouldStopTest(ctx context.Context, session *models.ExamSession) bool {
	responses, err := s.examSessionRepo.GetResponses(ctx, session.ID)
	if err != nil {
		return true
	}

	// 将 QuestionResponse 转换为 Question 切片
	var questions []*models.Question
	for _, resp := range responses {
		question, err := s.questionRepo.FindByID(ctx, resp.QuestionID)
		if err != nil {
			continue
		}
		questions = append(questions, question)
	}

	// 计算当前信息量
	info := s.calculateTestInformation(ctx, session.CurrentAbility, questions)
	se := utils.CalculateStandardError(info)

	return se < 0.3 || len(responses) >= 30
}
