package dto

import "time"

// PageQuery 分页查询参数
type PageQuery struct {
	Page     int    `form:"page" binding:"min=1" default:"1"`
	PageSize int    `form:"page_size" binding:"min=1,max=100" default:"20"`
	Sort     string `form:"sort"`
	Order    string `form:"order" binding:"oneof=asc desc"`
}

func (q *PageQuery) GetOffset() int {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 1 {
		q.PageSize = 20
	}
	return (q.Page - 1) * q.PageSize
}

func (q *PageQuery) GetLimit() int {
	if q.PageSize < 1 {
		q.PageSize = 20
	}
	if q.PageSize > 100 {
		q.PageSize = 100
	}
	return q.PageSize
}

// PageResponse 分页响应
type PageResponse struct {
	Data interface{} `json:"data"`
	Meta PageMeta    `json:"meta"`
}

// PageMeta 分页元数据
type PageMeta struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int   `json:"total_pages"`
}

// NewPageResponse 创建分页响应
func NewPageResponse(data interface{}, total int64, page, pageSize int) PageResponse {
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return PageResponse{
		Data: data,
		Meta: PageMeta{
			Total:      total,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		},
	}
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(code string, message string, details interface{}) ErrorResponse {
	return ErrorResponse{
		Code:      code,
		Message:   message,
		Details:   details,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// SuccessResponse 成功响应
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Message: message,
		Data:    data,
	}
}
