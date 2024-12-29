package middleware

import (
	"github.com/gin-gonic/gin"
)

// APIVersion API版本中间件
func APIVersion(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 添加API版本到响应头
		c.Header("X-API-Version", version)
		// 将版本信息添加到上下文
		c.Set("api_version", version)
		c.Next()
	}
}
