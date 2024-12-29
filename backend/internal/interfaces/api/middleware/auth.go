package middleware

import (
	"net/http"
	"strings"

	"irt-exam-system/backend/internal/domain/repositories"

	"github.com/gin-gonic/gin"
)

// Auth returns a middleware handler that validates JWT tokens
func Auth(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		token := parts[1]
		_ = token
		// TODO: 实现token验证逻辑
		// user, err := userRepo.FindByToken(c.Request.Context(), token)
		// if err != nil {
		//     c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		//     return
		// }
		// c.Set("user", user)

		c.Next()
	}
}
