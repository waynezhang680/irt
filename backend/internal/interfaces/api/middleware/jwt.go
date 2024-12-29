package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

// Claims JWT claims结构
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID uint, role string) (string, error) {
	// 从环境变量获取密钥
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		secretKey = []byte("your-256-bit-secret") // 默认密钥，生产环境应该使用环境变量
	}

	// 设置claims
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 24小时过期
			IssuedAt:  time.Now().Unix(),
			Issuer:    "irt-exam-system",
		},
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	// 从环境变量获取密钥
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		secretKey = []byte("your-256-bit-secret") // 默认密钥，生产环境应该使用环境变量
	}

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrExpiredToken
			}
		}
		return nil, ErrInvalidToken
	}

	// 验证token类型
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// RefreshToken 刷新JWT token
func RefreshToken(tokenString string) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil && err != ErrExpiredToken {
		return "", err
	}

	// 如果token未过期，且剩余时间超过12小时，不刷新
	if err != ErrExpiredToken {
		remaining := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())
		if remaining > 12*time.Hour {
			return tokenString, nil
		}
	}

	// 生成新token
	return GenerateToken(claims.UserID, claims.Role)
}
