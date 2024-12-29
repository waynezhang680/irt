package handlers

import (
	"net/http"
	"strconv"

	"irt-exam-system/backend/internal/application/services"
	"irt-exam-system/backend/internal/interfaces/api/dto"

	"github.com/gin-gonic/gin"
)

// AbilityHandler handles ability related requests
type AbilityHandler struct {
	abilityService services.AbilityService
}

// NewAbilityHandler creates a new ability handler
func NewAbilityHandler(abilityService services.AbilityService) *AbilityHandler {
	return &AbilityHandler{
		abilityService: abilityService,
	}
}

// List returns a list of abilities
func (h *AbilityHandler) List(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid user ID", err.Error()))
		return
	}

	abilities, err := h.abilityService.ListUserAbilities(c, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get abilities", err.Error()))
		return
	}

	c.JSON(http.StatusOK, abilities)
}

// GetBySubject returns ability by subject
func (h *AbilityHandler) GetBySubject(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid user ID", err.Error()))
		return
	}

	subjectID, err := strconv.ParseUint(c.Param("subject_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid subject ID", err.Error()))
		return
	}

	ability, err := h.abilityService.GetUserAbility(c, uint(userID), uint(subjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get ability", err.Error()))
		return
	}

	if ability == nil {
		c.JSON(http.StatusNotFound, dto.NewErrorResponse("404", "Ability not found", nil))
		return
	}

	c.JSON(http.StatusOK, ability)
}

// GetHistory returns ability estimation history
func (h *AbilityHandler) GetHistory(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid user ID", err.Error()))
		return
	}

	var query dto.PageQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("400", "Invalid query parameters", err.Error()))
		return
	}

	history, total, err := h.abilityService.GetUserEstimationHistory(c, uint(userID), query.GetOffset(), query.GetLimit())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get ability history", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewPageResponse(history, total, query.Page, query.PageSize))
}
