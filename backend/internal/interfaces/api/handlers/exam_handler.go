package handlers

import (
	"net/http"
	"strconv"

	"irt-exam-system/backend/internal/application/services"
	"irt-exam-system/backend/internal/interfaces/api/dto"

	"github.com/gin-gonic/gin"
)

// ExamHandler handles exam related requests
type ExamHandler struct {
	examService services.ExamService
}

// NewExamHandler creates a new exam handler
func NewExamHandler(examService services.ExamService) *ExamHandler {
	return &ExamHandler{
		examService: examService,
	}
}

// List returns a list of exams
func (h *ExamHandler) List(c *gin.Context) {
	var query dto.PageQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid query parameters", err.Error()))
		return
	}

	filters := make(map[string]interface{})
	exams, total, err := h.examService.ListExamPapers(c, filters, query.GetOffset(), query.GetLimit())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get exams", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewPageResponse(exams, total, query.Page, query.PageSize))
}

// GetByID returns an exam by ID
func (h *ExamHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid exam ID", err.Error()))
		return
	}

	exam, err := h.examService.GetExamPaper(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get exam", err.Error()))
		return
	}

	if exam == nil {
		c.JSON(http.StatusNotFound, dto.NewErrorResponse("404", "Exam not found", nil))
		return
	}

	c.JSON(http.StatusOK, exam)
}

// SubmitAnswer submits an answer for a question
func (h *ExamHandler) SubmitAnswer(c *gin.Context) {
	recordID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid record ID", err.Error()))
		return
	}

	var req struct {
		QuestionID uint   `json:"question_id" binding:"required"`
		Answer     string `json:"answer" binding:"required"`
		TimeSpent  int64  `json:"time_spent" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid request body", err.Error()))
		return
	}

	response, err := h.examService.SubmitAnswer(c, uint(recordID), req.QuestionID, req.Answer, req.TimeSpent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to submit answer", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

// SubmitExam submits the entire exam
func (h *ExamHandler) SubmitExam(c *gin.Context) {
	recordID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid record ID", err.Error()))
		return
	}

	var req struct {
		TotalTime  int64 `json:"total_time" binding:"required"`
		AutoSubmit bool  `json:"auto_submit"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid request body", err.Error()))
		return
	}

	record, err := h.examService.FinishExam(c, uint(recordID), req.TotalTime, req.AutoSubmit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to submit exam", err.Error()))
		return
	}

	c.JSON(http.StatusOK, record)
}
