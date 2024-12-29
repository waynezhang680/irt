package repositories

import (
	"context"
	"errors"

	"irt-exam-system/backend/internal/domain/repositories"
	"irt-exam-system/backend/models"

	"gorm.io/gorm"
)

type knowledgeRepository struct {
	db *gorm.DB
}

// NewKnowledgeRepository creates a new knowledge point repository instance
func NewKnowledgeRepository(db *gorm.DB) repositories.KnowledgePointRepository {
	return &knowledgeRepository{db: db}
}

// 基本操作实现
func (r *knowledgeRepository) Create(ctx context.Context, point *models.KnowledgePoint) error {
	return r.db.WithContext(ctx).Create(point).Error
}

func (r *knowledgeRepository) Update(ctx context.Context, point *models.KnowledgePoint) error {
	return r.db.WithContext(ctx).Save(point).Error
}

func (r *knowledgeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.KnowledgePoint{}, id).Error
}

func (r *knowledgeRepository) FindByID(ctx context.Context, id uint) (*models.KnowledgePoint, error) {
	var point models.KnowledgePoint
	err := r.db.WithContext(ctx).First(&point, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &point, nil
}

// 查询操作实现
func (r *knowledgeRepository) ListBySubject(ctx context.Context, subjectID uint) ([]*models.KnowledgePoint, error) {
	var points []*models.KnowledgePoint
	err := r.db.WithContext(ctx).Where("subject_id = ?", subjectID).Find(&points).Error
	return points, err
}

func (r *knowledgeRepository) ListByParent(ctx context.Context, parentID uint) ([]*models.KnowledgePoint, error) {
	var points []*models.KnowledgePoint
	err := r.db.WithContext(ctx).Where("parent_id = ?", parentID).Find(&points).Error
	return points, err
}

func (r *knowledgeRepository) GetFullPath(ctx context.Context, id uint) ([]*models.KnowledgePoint, error) {
	var path []*models.KnowledgePoint
	current, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if current == nil {
		return path, nil
	}

	path = append(path, current)
	for current.ParentID != nil {
		current, err = r.FindByID(ctx, *current.ParentID)
		if err != nil {
			return nil, err
		}
		if current == nil {
			break
		}
		path = append([]*models.KnowledgePoint{current}, path...)
	}
	return path, nil
}

func (r *knowledgeRepository) Search(ctx context.Context, keyword string) ([]*models.KnowledgePoint, error) {
	var points []*models.KnowledgePoint
	err := r.db.WithContext(ctx).Where("name LIKE ? OR description LIKE ?",
		"%"+keyword+"%", "%"+keyword+"%").Find(&points).Error
	return points, err
}

// 题目关联操作实现
func (r *knowledgeRepository) ListQuestions(ctx context.Context, knowledgePointID uint, offset, limit int) ([]*models.Question, int64, error) {
	var questions []*models.Question
	var total int64

	subQuery := r.db.Table("question_knowledge_points").Select("question_id").
		Where("knowledge_point_id = ?", knowledgePointID)

	err := r.db.WithContext(ctx).Model(&models.Question{}).
		Where("id IN (?)", subQuery).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("id IN (?)", subQuery).
		Offset(offset).Limit(limit).Find(&questions).Error
	if err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

func (r *knowledgeRepository) AddQuestion(ctx context.Context, knowledgePointID, questionID uint) error {
	relation := &models.QuestionKnowledgePoint{
		QuestionID:       questionID,
		KnowledgePointID: knowledgePointID,
	}
	return r.db.WithContext(ctx).Create(relation).Error
}

func (r *knowledgeRepository) RemoveQuestion(ctx context.Context, knowledgePointID, questionID uint) error {
	return r.db.WithContext(ctx).Where("knowledge_point_id = ? AND question_id = ?",
		knowledgePointID, questionID).Delete(&models.QuestionKnowledgePoint{}).Error
}

// 树形结构操作实现
func (r *knowledgeRepository) MoveNode(ctx context.Context, id, newParentID uint) error {
	return r.db.WithContext(ctx).Model(&models.KnowledgePoint{}).
		Where("id = ?", id).Update("parent_id", newParentID).Error
}

func (r *knowledgeRepository) GetChildren(ctx context.Context, id uint) ([]*models.KnowledgePoint, error) {
	var children []*models.KnowledgePoint
	err := r.db.WithContext(ctx).Where("parent_id = ?", id).Find(&children).Error
	return children, err
}

func (r *knowledgeRepository) GetDescendants(ctx context.Context, id uint) ([]*models.KnowledgePoint, error) {
	var descendants []*models.KnowledgePoint
	var stack []uint
	stack = append(stack, id)

	for len(stack) > 0 {
		currentID := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		children, err := r.GetChildren(ctx, currentID)
		if err != nil {
			return nil, err
		}

		for _, child := range children {
			descendants = append(descendants, child)
			stack = append(stack, child.ID)
		}
	}

	return descendants, nil
}

func (r *knowledgeRepository) GetAncestors(ctx context.Context, id uint) ([]*models.KnowledgePoint, error) {
	path, err := r.GetFullPath(ctx, id)
	if err != nil {
		return nil, err
	}
	// 移除当前节点，只返回祖先节点
	if len(path) > 0 {
		path = path[:len(path)-1]
	}
	return path, nil
}
