package models

// UserRegisterRequest 用户注册请求
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// UserResponse 用户信息响应
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// ExamSessionResponse 考试会话响应
type ExamSessionResponse struct {
	ID             uint    `json:"id"`
	StartTime      string  `json:"start_time"`
	CurrentAbility float64 `json:"current_ability"`
}

// QuestionDTO 试题响应
type QuestionDTO struct {
	ID      uint     `json:"id"`
	Content string   `json:"content"`
	Options []string `json:"options"`
}

// AnswerRequest 答案请求
type AnswerRequest struct {
	QuestionID uint   `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
	TimeSpent  int    `json:"time_spent" binding:"required"`
}

// AnswerResponse 答案响应
type AnswerResponse struct {
	IsCorrect      bool    `json:"is_correct"`
	CurrentAbility float64 `json:"current_ability"`
	NextDifficulty float64 `json:"next_difficulty"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
