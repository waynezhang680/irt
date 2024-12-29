package services

import (
	"context"
	"time"

	"irt-exam-system/backend/internal/domain/models"
	"irt-exam-system/backend/internal/domain/repositories"
)

// ExamService 考试服务接口
type ExamService interface {
	CreateExamPaper(ctx context.Context, paper *models.ExamPaper) error
	UpdateExamPaper(ctx context.Context, paper *models.ExamPaper) error
	DeleteExamPaper(ctx context.Context, id uint) error
	GetExamPaper(ctx context.Context, id uint) (*models.ExamPaper, error)
	ListExamPapers(ctx context.Context, filters map[string]interface{}, offset, limit int) ([]*models.ExamPaper, int64, error)
	StartExam(ctx context.Context, userID, paperID uint) (*models.ExamRecord, error)
	SubmitAnswer(ctx context.Context, recordID, questionID uint, answer string, timeSpent int64) (*models.ExamResponse, error)
	FinishExam(ctx context.Context, recordID uint, totalTime int64, autoSubmit bool) (*models.ExamRecord, error)
	GetExamRecord(ctx context.Context, recordID uint) (*models.ExamRecord, error)
	ListUserExams(ctx context.Context, userID uint, offset, limit int) ([]*models.ExamRecord, int64, error)
	GetExamResult(ctx context.Context, recordID uint) (*models.ExamResult, error)
	GetQuestionAnalysis(ctx context.Context, recordID uint) ([]*QuestionAnalysis, error)
}

// ExamResult 考试结果
type ExamResult struct {
	ExamID         uint      `json:"exam_id"`
	Title          string    `json:"title"`
	Score          float64   `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	CorrectCount   int       `json:"correct_count"`
	IncorrectCount int       `json:"incorrect_count"`
	TimeTaken      string    `json:"time_taken"`
	SubmitTime     time.Time `json:"submit_time"`
	Analysis       struct {
		AbilityEstimate    float64  `json:"ability_estimate"`
		ConfidenceInterval string   `json:"confidence_interval"`
		PerformanceLevel   string   `json:"performance_level"`
		Recommendations    []string `json:"recommendations"`
	} `json:"analysis"`
}

// QuestionAnalysis 题目分析
type QuestionAnalysis struct {
	ID         uint    `json:"id"`
	IsCorrect  bool    `json:"is_correct"`
	TimeSpent  int64   `json:"time_spent"`
	Score      float64 `json:"score"`
	Difficulty float64 `json:"difficulty"`
}

// NewExamService creates a new exam service instance
func NewExamService(examRepo repositories.ExamRepository, questionRepo repositories.QuestionRepository) ExamService {
	return &examService{
		examRepo:     examRepo,
		questionRepo: questionRepo,
	}
}

type examService struct {
	examRepo     repositories.ExamRepository
	questionRepo repositories.QuestionRepository
}

// CreateExamPaper implements ExamService
func (s *examService) CreateExamPaper(ctx context.Context, paper *models.ExamPaper) error {
	return s.examRepo.CreatePaper(ctx, paper)
}

// UpdateExamPaper implements ExamService
func (s *examService) UpdateExamPaper(ctx context.Context, paper *models.ExamPaper) error {
	return s.examRepo.UpdatePaper(ctx, paper)
}

// DeleteExamPaper implements ExamService
func (s *examService) DeleteExamPaper(ctx context.Context, id uint) error {
	return s.examRepo.DeletePaper(ctx, id)
}

// GetExamPaper implements ExamService
func (s *examService) GetExamPaper(ctx context.Context, id uint) (*models.ExamPaper, error) {
	return s.examRepo.FindPaperByID(ctx, id)
}

// ListExamPapers implements ExamService
func (s *examService) ListExamPapers(ctx context.Context, filters map[string]interface{}, offset, limit int) ([]*models.ExamPaper, int64, error) {
	return s.examRepo.ListPapers(ctx, offset, limit)
}

// StartExam implements ExamService
func (s *examService) StartExam(ctx context.Context, userID, paperID uint) (*models.ExamRecord, error) {
	record := &models.ExamRecord{
		UserID:    userID,
		ExamID:    paperID,
		StartTime: time.Now(),
		Status:    "in_progress",
	}
	err := s.examRepo.CreateRecord(ctx, record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// SubmitAnswer implements ExamService
func (s *examService) SubmitAnswer(ctx context.Context, recordID, questionID uint, answer string, timeSpent int64) (*models.ExamResponse, error) {
	// 获取题目信息
	question, err := s.questionRepo.FindByID(ctx, questionID)
	if err != nil {
		return nil, err
	}

	// 判断答案是否正确
	isCorrect := question.Answer == answer
	score := 0.0
	if isCorrect {
		score = question.Score
	}

	// 创建答题记录
	response := &models.ExamResponse{
		RecordID:     recordID,
		QuestionID:   questionID,
		UserAnswer:   answer,
		IsCorrect:    isCorrect,
		ResponseTime: timeSpent,
		Score:        score,
	}

	err = s.examRepo.CreateResponse(ctx, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// FinishExam implements ExamService
func (s *examService) FinishExam(ctx context.Context, recordID uint, totalTime int64, autoSubmit bool) (*models.ExamRecord, error) {
	record, err := s.examRepo.FindRecordByID(ctx, recordID)
	if err != nil {
		return nil, err
	}

	record.EndTime = time.Now()
	record.Status = "completed"

	// 计算总分
	responses, err := s.examRepo.ListRecordResponses(ctx, recordID)
	if err != nil {
		return nil, err
	}

	var totalScore float64
	for _, response := range responses {
		totalScore += response.Score
	}
	record.Score = totalScore

	err = s.examRepo.UpdateRecord(ctx, record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// GetExamRecord implements ExamService
func (s *examService) GetExamRecord(ctx context.Context, recordID uint) (*models.ExamRecord, error) {
	return s.examRepo.FindRecordByID(ctx, recordID)
}

// ListUserExams implements ExamService
func (s *examService) ListUserExams(ctx context.Context, userID uint, offset, limit int) ([]*models.ExamRecord, int64, error) {
	return s.examRepo.ListUserRecords(ctx, userID, offset, limit)
}

// GetExamResult implements ExamService
func (s *examService) GetExamResult(ctx context.Context, recordID uint) (*models.ExamResult, error) {
	// 获取考试记录
	record, err := s.examRepo.FindRecordByID(ctx, recordID)
	if err != nil {
		return nil, err
	}

	// 获取答题记录
	responses, err := s.examRepo.ListRecordResponses(ctx, recordID)
	if err != nil {
		return nil, err
	}

	// 计算总分
	var totalScore float64
	for _, response := range responses {
		totalScore += response.Score
	}

	result := &models.ExamResult{
		ExamID:         record.ExamPaperID,
		Title:          record.ExamPaper.Title,
		Score:          totalScore,
		TotalQuestions: len(responses),
		SubmitTime:     record.EndTime,
	}

	for _, response := range responses {
		if response.IsCorrect {
			result.CorrectCount++
		} else {
			result.IncorrectCount++
		}
	}

	result.TimeTaken = record.EndTime.Sub(record.StartTime).String()

	return result, nil
}

// GetQuestionAnalysis implements ExamService
func (s *examService) GetQuestionAnalysis(ctx context.Context, recordID uint) ([]*QuestionAnalysis, error) {
	responses, err := s.examRepo.ListRecordResponses(ctx, recordID)
	if err != nil {
		return nil, err
	}

	var analysis []*QuestionAnalysis
	for _, response := range responses {
		question, err := s.questionRepo.FindByID(ctx, response.QuestionID)
		if err != nil {
			continue
		}

		analysis = append(analysis, &QuestionAnalysis{
			ID:         response.QuestionID,
			IsCorrect:  response.IsCorrect,
			TimeSpent:  response.ResponseTime,
			Score:      response.Score,
			Difficulty: question.Difficulty,
		})
	}

	return analysis, nil
}
