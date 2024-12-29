package handlers

import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "backend/internal/application/services"
    "backend/internal/domain/models"
    "backend/internal/infrastructure/persistence"
    "backend/pkg/errors"
)

type AuthHandler struct {
    AuthService *services.AuthService
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
       
} 