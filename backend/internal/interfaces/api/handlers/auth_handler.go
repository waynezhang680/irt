package handlers

import (
	"net/http"

	"irt-exam-system/backend/internal/application/services"
	"irt-exam-system/backend/internal/interfaces/api/dto"
	"irt-exam-system/backend/models"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证处理器
type AuthHandler struct {
	userService services.UserService
}

// NewAuthHandler 创建认证处理器
func NewAuthHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// Login godoc
// @Summary 用户登录
// @Description 用户登录并获取JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "登录信息"
// @Success 200 {object} dto.LoginResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("200", "Invalid request body", err.Error()))
		return
	}

	token, err := h.userService.Login(c, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("401", "Login failed", err.Error()))
		return
	}

	user, err := h.userService.GetUserByEmail(c, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to get user info", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
		User: dto.UserInfo{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			RoleType:    string(user.RoleType),
			Permissions: user.Permissions,
		},
		ExpiresIn: 86400, // 24小时
	})
}

// Register godoc
// @Summary 用户注册
// @Description 注册新用户
// @Tags auth
// @Accept json
// @Produce json
// @Param register body dto.RegisterRequest true "注册信息"
// @Success 200 {object} dto.RegisterResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("200", "Invalid request body", err.Error()))
		return
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		RoleType: models.RoleStudent,
	}

	err := h.userService.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("500", "Failed to create user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.RegisterResponse{
		Message: "User registered successfully",
		UserID:  user.ID,
	})
}

// RefreshToken godoc
// @Summary 刷新token
// @Description 使用旧token获取新token
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body dto.RefreshTokenRequest true "刷新token请求"
// @Success 200 {object} dto.RefreshTokenResponse
// @Router /auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("200", "Invalid request body", err.Error()))
		return
	}

	newToken, err := h.userService.RefreshToken(c, req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("401", "Failed to refresh token", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		Token:     newToken,
		ExpiresIn: 86400, // 24小时
	})
}
