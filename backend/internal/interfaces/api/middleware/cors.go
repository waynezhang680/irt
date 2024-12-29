package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS returns a middleware handler that adds CORS headers to the response
func CORS() gin.HandlerFunc {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	allowedMethods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")
	allowedHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")
	maxAge, _ := time.ParseDuration(os.Getenv("CORS_MAX_AGE") + "s")

	config := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     allowedMethods,
		AllowHeaders:     allowedHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	}

	return cors.New(config)
}
