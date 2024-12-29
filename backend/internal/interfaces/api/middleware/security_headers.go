package middleware

import "github.com/gin-gonic/gin"

// SecurityHeaders 添加安全相关的HTTP头
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 防止XSS攻击
		c.Header("X-XSS-Protection", "1; mode=block")
		// 防止点击劫持
		c.Header("X-Frame-Options", "DENY")
		// 防止MIME类型嗅探
		c.Header("X-Content-Type-Options", "nosniff")
		// 启用HSTS
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// 内容安全策略
		c.Header("Content-Security-Policy", "default-src 'self'")
		// 引用策略
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		// 特性策略
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		c.Next()
	}
}
