package handlers

import (
	"net/http"
	"strconv"

	"irt-exam-system/backend/internal/application/services"
	"irt-exam-system/backend/internal/interfaces/api/dto"

	"github.com/gin-gonic/gin"
)

// KnowledgeHandler handles knowledge point related requests
type KnowledgeHandler struct {
	knowledgeService services.KnowledgeService
}

// NewKnowledgeHandler creates a new knowledge handler
func NewKnowledgeHandler(knowledgeService services.KnowledgeService) *KnowledgeHandler {
	return &KnowledgeHandler{
		knowledgeService: knowledgeService,
	}
}

// List returns a list of knowledge points
func (h *KnowledgeHandler) List(c *gin.Context) {
	var query dto.PageQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid query parameters", err.Error()))
		return
	}

	points, total, err := h.knowledgeService.ListKnowledgePoints(c, query.GetOffset(), query.GetLimit())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get knowledge points", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewPageResponse(points, total, query.Page, query.PageSize))
}

// GetByID returns a knowledge point by ID
func (h *KnowledgeHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid knowledge point ID", err.Error()))
		return
	}

	point, err := h.knowledgeService.GetKnowledgePoint(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get knowledge point", err.Error()))
		return
	}

	if point == nil {
		c.JSON(http.StatusNotFound, dto.NewErrorResponse("404", "Knowledge point not found", nil))
		return
	}

	c.JSON(http.StatusOK, point)
}

// GetProgress returns the progress of a knowledge point
func (h *KnowledgeHandler) GetProgress(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid knowledge point ID", err.Error()))
		return
	}

	questions, err := h.knowledgeService.GetKnowledgePointQuestions(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get knowledge point questions", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"knowledge_point_id": id,
		"total_questions":    len(questions),
		"questions":          questions,
	})
}
