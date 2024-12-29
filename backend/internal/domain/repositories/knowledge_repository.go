package repositories

import (
	"context"

	"irt-exam-system/backend/models"
)

// KnowledgePointRepository 知识点仓储接口
type KnowledgePointRepository interface {
	// 基本操作
	Create(ctx context.Context, point *models.KnowledgePoint) error
	Update(ctx context.Context, point *models.KnowledgePoint) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*models.KnowledgePoint, error)

	// 查询操作
	ListBySubject(ctx context.Context, subjectID uint) ([]*models.KnowledgePoint, error)
	ListByParent(ctx context.Context, parentID uint) ([]*models.KnowledgePoint, error)
	GetFullPath(ctx context.Context, id uint) ([]*models.KnowledgePoint, error)
	Search(ctx context.Context, keyword string) ([]*models.KnowledgePoint, error)

	// 题目关联操作
	ListQuestions(ctx context.Context, knowledgePointID uint, offset, limit int) ([]*models.Question, int64, error)
	AddQuestion(ctx context.Context, knowledgePointID, questionID uint) error
	RemoveQuestion(ctx context.Context, knowledgePointID, questionID uint) error

	// 树形结构操作
	MoveNode(ctx context.Context, id, newParentID uint) error
	GetChildren(ctx context.Context, id uint) ([]*models.KnowledgePoint, error)
	GetDescendants(ctx context.Context, id uint) ([]*models.KnowledgePoint, error)
	GetAncestors(ctx context.Context, id uint) ([]*models.KnowledgePoint, error)
}
