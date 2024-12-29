package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// 预定义错误
var (
	ErrInvalidInput   = &AppError{Code: 400, Message: "Invalid input"}
	ErrUnauthorized   = &AppError{Code: 401, Message: "Unauthorized"}
	ErrForbidden      = &AppError{Code: 403, Message: "Forbidden"}
	ErrNotFound       = &AppError{Code: 404, Message: "Resource not found"}
	ErrInternalServer = &AppError{Code: 500, Message: "Internal server error"}
)
