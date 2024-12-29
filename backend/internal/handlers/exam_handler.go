package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary 开始考试
// @Description 开始一个新的考试会话
// @Tags 考试
// @Accept json
// @Produce json
// @Security Bearer
// @Success 201 {object} models.ExamSessionResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /exam/start [post]
func StartExam(c *gin.Context) {
	// TODO: 实现开始考试逻辑
}

// @Summary 获取下一题
// @Description 根据IRT算法获取下一道试题
// @Tags 考试
// @Accept json
// @Produce json
// @Security Bearer
// @Param session_id path int true "考试会话ID"
// @Success 200 {object} models.QuestionResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /exam/{session_id}/next [get]
func GetNextQuestion(c *gin.Context) {
	// TODO: 实现获取下一题逻辑
}

// @Summary 提交答案
// @Description 提交试题答案并获取评估结果
// @Tags 考试
// @Accept json
// @Produce json
// @Security Bearer
// @Param session_id path int true "考试会话ID"
// @Param answer body models.AnswerRequest true "答案信息"
// @Success 200 {object} models.AnswerResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /exam/{session_id}/answer [post]
func SubmitAnswer(c *gin.Context) {
	// TODO: 实现提交答案逻辑
}
