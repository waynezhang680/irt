package dto

import (
	"time"

	"irt-exam-system/backend/internal/application/services"
	"irt-exam-system/backend/models"
)

// ExamListQuery 考试列表查询参数
type ExamListQuery struct {
	Status    string `form:"status"`
	SubjectID uint   `form:"subject_id"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	PageQuery
}

func (q *ExamListQuery) ToFilters() map[string]interface{} {
	filters := make(map[string]interface{})
	if q.Status != "" {
		filters["status"] = q.Status
	}
	if q.SubjectID > 0 {
		filters["subject_id"] = q.SubjectID
	}
	if q.StartDate != "" {
		filters["start_date"] = q.StartDate
	}
	if q.EndDate != "" {
		filters["end_date"] = q.EndDate
	}
	return filters
}

// ExamPaperResponse 考试列表响应
type ExamPaperResponse struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Duration       int64     `json:"duration"`
	TotalQuestions int       `json:"total_questions"`
	TotalScore     float64   `json:"total_score"`
	PassScore      float64   `json:"pass_score"`
	Status         string    `json:"status"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	SubjectID      uint      `json:"subject_id"`
	SubjectName    string    `json:"subject_name"`
}

// ExamDetailResponse 考试详情响应
type ExamDetailResponse struct {
	ID           uint             `json:"id"`
	Title        string           `json:"title"`
	Duration     int64            `json:"duration"`
	Questions    []QuestionDetail `json:"questions"`
	TotalScore   float64          `json:"total_score"`
	PassScore    float64          `json:"pass_score"`
	StartTime    time.Time        `json:"start_time"`
	EndTime      time.Time        `json:"end_time"`
	Instructions string           `json:"instructions"`
	SubjectID    uint             `json:"subject_id"`
	SubjectName  string           `json:"subject_name"`
}

// QuestionDetail 题目详情
type QuestionDetail struct {
	ID      uint     `json:"id"`
	Type    string   `json:"type"`
	Content string   `json:"content"`
	Options []Option `json:"options"`
	Score   float64  `json:"score"`
	Order   int64    `json:"order"`
	Section string   `json:"section,omitempty"`
}

// Option 选项
type Option struct {
	Label   string `json:"label"`
	Content string `json:"content"`
}

// SubmitAnswerRequest 提交答案请求
type SubmitAnswerRequest struct {
	RecordID   uint   `json:"record_id" binding:"required"`
	QuestionID uint   `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
	TimeSpent  int64  `json:"time_spent" binding:"required"`
}

// SubmitAnswerResponse 提交答案响应
type SubmitAnswerResponse struct {
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	Score     float64 `json:"score,omitempty"`
	IsCorrect bool    `json:"is_correct,omitempty"`
}

// SubmitExamRequest 提交考试请求
type SubmitExamRequest struct {
	RecordID   uint                  `json:"record_id" binding:"required"`
	Answers    []SubmitAnswerRequest `json:"answers"`
	TotalTime  int64                 `json:"total_time" binding:"required"`
	AutoSubmit bool                  `json:"auto_submit"`
}

// SubmitExamResponse 提交考试响应
type SubmitExamResponse struct {
	Status          string    `json:"status"`
	Message         string    `json:"message"`
	ExamID          uint      `json:"exam_id"`
	Score           float64   `json:"score"`
	Pass            bool      `json:"pass"`
	CompleteTime    time.Time `json:"complete_time"`
	ReviewAvailable bool      `json:"review_available"`
}

// ExamResultResponse 考试结果响应
type ExamResultResponse struct {
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
	QuestionDetails []QuestionAnalysisResponse `json:"question_details"`
}

// QuestionAnalysisResponse 题目分析响应
type QuestionAnalysisResponse struct {
	ID         uint    `json:"id"`
	IsCorrect  bool    `json:"is_correct"`
	TimeSpent  int64   `json:"time_spent"`
	Score      float64 `json:"score"`
	Difficulty float64 `json:"difficulty"`
}

// Conversion functions
func ToExamPaperResponses(papers []*models.ExamPaper) []ExamPaperResponse {
	responses := make([]ExamPaperResponse, len(papers))
	for i, paper := range papers {
		responses[i] = ExamPaperResponse{
			ID:             paper.ID,
			Title:          paper.Title,
			Duration:       paper.TimeLimit,
			TotalQuestions: len(paper.Questions),
			TotalScore:     paper.TotalScore,
			PassScore:      paper.PassScore,
			Status:         paper.Status,
			StartTime:      paper.StartTime,
			EndTime:        paper.EndTime,
			SubjectID:      paper.SubjectID,
			SubjectName:    paper.Subject.Name,
		}
	}
	return responses
}

func ToExamDetailResponse(paper *models.ExamPaper) ExamDetailResponse {
	questions := make([]QuestionDetail, len(paper.Questions))
	for i, q := range paper.Questions {
		options := make([]Option, len(q.Question.Options))
		for j, opt := range q.Question.Options {
			options[j] = Option{
				Label:   opt.Label,
				Content: opt.Content,
			}
		}
		questions[i] = QuestionDetail{
			ID:      q.QuestionID,
			Type:    q.Question.Type,
			Content: q.Question.Content,
			Options: options,
			Score:   q.Score,
			Order:   q.Order,
		}
	}
	return ExamDetailResponse{
		ID:          paper.ID,
		Title:       paper.Title,
		Duration:    paper.TimeLimit,
		Questions:   questions,
		TotalScore:  paper.TotalScore,
		PassScore:   paper.PassScore,
		StartTime:   paper.StartTime,
		EndTime:     paper.EndTime,
		SubjectID:   paper.SubjectID,
		SubjectName: paper.Subject.Name,
	}
}

func ToSubmitAnswerResponse(response *models.ExamResponse) SubmitAnswerResponse {
	return SubmitAnswerResponse{
		Status:    "success",
		Message:   "Answer submitted successfully",
		Score:     response.Score,
		IsCorrect: response.IsCorrect,
	}
}

func ToSubmitExamResponse(record *models.ExamRecord) SubmitExamResponse {
	return SubmitExamResponse{
		Status:          "success",
		Message:         "Exam submitted successfully",
		ExamID:          record.ExamPaperID,
		Score:           record.Score,
		Pass:            record.Score >= record.ExamPaper.PassScore,
		CompleteTime:    record.EndTime,
		ReviewAvailable: true,
	}
}

func ToExamResultResponse(result *services.ExamResult, analysis []*services.QuestionAnalysis) ExamResultResponse {
	resp := ExamResultResponse{
		ExamID:         result.ExamID,
		Title:          result.Title,
		Score:          result.Score,
		TotalQuestions: result.TotalQuestions,
		CorrectCount:   result.CorrectCount,
		IncorrectCount: result.IncorrectCount,
		TimeTaken:      result.TimeTaken,
		SubmitTime:     result.SubmitTime,
		Analysis:       result.Analysis,
	}

	resp.QuestionDetails = make([]QuestionAnalysisResponse, len(analysis))
	for i, a := range analysis {
		resp.QuestionDetails[i] = QuestionAnalysisResponse{
			ID:         a.ID,
			IsCorrect:  a.IsCorrect,
			TimeSpent:  a.TimeSpent,
			Score:      a.Score,
			Difficulty: a.Difficulty,
		}
	}
	return resp
}
