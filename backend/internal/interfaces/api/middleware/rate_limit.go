package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"irt-exam-system/backend/internal/interfaces/api/dto"

	"github.com/gin-gonic/gin"
)

// GetUserID 从上下文中获取用户ID
func GetUserID(c *gin.Context) uint {
	if user, exists := c.Get("user"); exists {
		if userMap, ok := user.(map[string]interface{}); ok {
			if id, exists := userMap["id"]; exists {
				if idFloat, ok := id.(float64); ok {
					return uint(idFloat)
				}
			}
		}
	}
	return 0
}

type rateLimiter struct {
	tokens     int
	lastRefill time.Time
	mu         sync.Mutex
}

var (
	ipLimiters   = make(map[string]*rateLimiter)
	userLimiters = make(map[uint]*rateLimiter)
	limiterMu    sync.RWMutex
)

// RateLimit 速率限制中间件
func RateLimit(limit int, per string) gin.HandlerFunc {
	duration := parseDuration(per)
	return func(c *gin.Context) {
		var key string
		var limiter *rateLimiter

		// 优先使用用户ID作为限制key
		if userID := GetUserID(c); userID > 0 {
			key = fmt.Sprintf("user:%d", userID)
			limiter = getUserLimiter(userID, limit, duration)
		} else {
			// 否则使用IP地址
			key = c.ClientIP()
			limiter = getIPLimiter(key, limit, duration)
		}

		if !checkRateLimit(limiter, duration) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, dto.ErrorResponse{
				Code:    "429",
				Message: "Rate limit exceeded",
			})
			return
		}

		// 设置速率限制响应头
		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limit))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.tokens))
		c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(duration).Unix()))

		c.Next()
	}
}

func getUserLimiter(userID uint, limit int, duration time.Duration) *rateLimiter {
	limiterMu.Lock()
	defer limiterMu.Unlock()

	if l, exists := userLimiters[userID]; exists {
		return l
	}

	l := &rateLimiter{
		tokens:     limit,
		lastRefill: time.Now(),
	}
	userLimiters[userID] = l
	return l
}

func getIPLimiter(ip string, limit int, duration time.Duration) *rateLimiter {
	limiterMu.Lock()
	defer limiterMu.Unlock()

	if l, exists := ipLimiters[ip]; exists {
		return l
	}

	l := &rateLimiter{
		tokens:     limit,
		lastRefill: time.Now(),
	}
	ipLimiters[ip] = l
	return l
}

func checkRateLimit(l *rateLimiter, duration time.Duration) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.lastRefill)

	// 如果已经过了一个时间周期，重置令牌
	if elapsed >= duration {
		l.tokens = 1
		l.lastRefill = now
		return true
	}

	// 如果还有令牌，消耗一个
	if l.tokens > 0 {
		l.tokens--
		return true
	}

	return false
}

func parseDuration(per string) time.Duration {
	switch per {
	case "second":
		return time.Second
	case "minute":
		return time.Minute
	case "hour":
		return time.Hour
	default:
		return time.Hour
	}
}
