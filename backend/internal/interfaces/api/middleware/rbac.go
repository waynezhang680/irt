package middleware

import (
	"net/http"

	"irt-exam-system/backend/internal/interfaces/api/dto"
	"irt-exam-system/backend/models"

	"github.com/gin-gonic/gin"
)

// RequireRole 检查用户角色
func RequireRole(requiredRole models.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewErrorResponse("401", "Unauthorized", "User not found"))
			return
		}

		if userMap, ok := user.(map[string]interface{}); ok {
			if roleType, exists := userMap["role_type"]; exists {
				if roleStr, ok := roleType.(string); ok {
					if models.RoleType(roleStr) == requiredRole {
						c.Next()
						return
					}
				}
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, dto.NewErrorResponse("403", "Forbidden", "Insufficient privileges"))
	}
}
