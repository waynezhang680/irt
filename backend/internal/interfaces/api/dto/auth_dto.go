package dto

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Captcha  string `json:"captcha,omitempty"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string   `json:"token"`
	User      UserInfo `json:"user"`
	ExpiresIn int      `json:"expires_in"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Captcha  string `json:"captcha" binding:"required"`
}

// RegisterResponse 注册响应
type RegisterResponse struct {
	Message        string `json:"message"`
	UserID         uint   `json:"user_id"`
	ActivationLink string `json:"activation_link,omitempty"`
}

// RefreshTokenRequest 刷新token请求
type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// RefreshTokenResponse 刷新token响应
type RefreshTokenResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID          uint     `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	RoleType    string   `json:"role_type"`
	Permissions []string `json:"permissions"`
}
