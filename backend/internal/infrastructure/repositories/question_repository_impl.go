package repositories

import (
	"context"

	"irt-exam-system/backend/internal/domain/models"
	"irt-exam-system/backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type QuestionRepositoryImpl struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) repositories.QuestionRepository {
	return &QuestionRepositoryImpl{db: db}
}

// Create implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) Create(ctx context.Context, question *models.Question) error {
	return r.db.WithContext(ctx).Create(question).Error
}

// Update implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) Update(ctx context.Context, question *models.Question) error {
	return r.db.WithContext(ctx).Save(question).Error
}

// Delete implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Question{}, id).Error
}

// FindByID implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) FindByID(ctx context.Context, id uint) (*models.Question, error) {
	var question models.Question
	err := r.db.WithContext(ctx).First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

// ListBySubject implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) ListBySubject(ctx context.Context, subjectID uint, offset, limit int) ([]*models.Question, int64, error) {
	var questions []*models.Question
	var total int64

	err := r.db.WithContext(ctx).Model(&models.Question{}).Where("subject_id = ?", subjectID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("subject_id = ?", subjectID).Offset(offset).Limit(limit).Find(&questions).Error
	if err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

// ListByKnowledgePoint implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) ListByKnowledgePoint(ctx context.Context, knowledgePointID uint, offset, limit int) ([]*models.Question, int64, error) {
	var questions []*models.Question
	var total int64

	query := r.db.WithContext(ctx).
		Joins("JOIN question_knowledge_points ON questions.id = question_knowledge_points.question_id").
		Where("question_knowledge_points.knowledge_point_id = ?", knowledgePointID)

	err := query.Model(&models.Question{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Find(&questions).Error
	if err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

// ListByType implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) ListByType(ctx context.Context, questionType string, offset, limit int) ([]*models.Question, int64, error) {
	var questions []*models.Question
	var total int64

	err := r.db.WithContext(ctx).Model(&models.Question{}).Where("type = ?", questionType).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("type = ?", questionType).Offset(offset).Limit(limit).Find(&questions).Error
	if err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

// Search implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) Search(ctx context.Context, keyword string, offset, limit int) ([]*models.Question, int64, error) {
	var questions []*models.Question
	var total int64

	query := r.db.WithContext(ctx).Where("content LIKE ? OR analysis LIKE ?", "%"+keyword+"%", "%"+keyword+"%")

	err := query.Model(&models.Question{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Find(&questions).Error
	if err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

// ListByExamPaper implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) ListByExamPaper(ctx context.Context, examPaperID uint) ([]*models.Question, error) {
	var questions []*models.Question
	err := r.db.WithContext(ctx).
		Joins("JOIN exam_paper_questions ON questions.id = exam_paper_questions.question_id").
		Where("exam_paper_questions.exam_paper_id = ?", examPaperID).
		Find(&questions).Error
	return questions, err
}

// CreateOption implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) CreateOption(ctx context.Context, option *models.QuestionOption) error {
	return r.db.WithContext(ctx).Create(option).Error
}

// UpdateOption implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) UpdateOption(ctx context.Context, option *models.QuestionOption) error {
	return r.db.WithContext(ctx).Save(option).Error
}

// DeleteOption implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) DeleteOption(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.QuestionOption{}, id).Error
}

// ListOptions implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) ListOptions(ctx context.Context, questionID uint) ([]*models.QuestionOption, error) {
	var options []*models.QuestionOption
	err := r.db.WithContext(ctx).Where("question_id = ?", questionID).Find(&options).Error
	return options, err
}

// AddKnowledgePoint implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) AddKnowledgePoint(ctx context.Context, questionID, knowledgePointID uint) error {
	relation := &models.QuestionKnowledgePoint{
		QuestionID:       questionID,
		KnowledgePointID: knowledgePointID,
	}
	return r.db.WithContext(ctx).Create(relation).Error
}

// RemoveKnowledgePoint implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) RemoveKnowledgePoint(ctx context.Context, questionID, knowledgePointID uint) error {
	return r.db.WithContext(ctx).Where("question_id = ? AND knowledge_point_id = ?", questionID, knowledgePointID).
		Delete(&models.QuestionKnowledgePoint{}).Error
}

// ListKnowledgePoints implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) ListKnowledgePoints(ctx context.Context, questionID uint) ([]*models.KnowledgePoint, error) {
	var points []*models.KnowledgePoint
	err := r.db.WithContext(ctx).
		Joins("JOIN question_knowledge_points ON knowledge_points.id = question_knowledge_points.knowledge_point_id").
		Where("question_knowledge_points.question_id = ?", questionID).
		Find(&points).Error
	return points, err
}

// UpdateParameters implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) UpdateParameters(ctx context.Context, params *models.QuestionParameter) error {
	return r.db.WithContext(ctx).Save(params).Error
}

// GetParameters implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) GetParameters(ctx context.Context, questionID uint) (*models.QuestionParameter, error) {
	var params models.QuestionParameter
	err := r.db.WithContext(ctx).Where("question_id = ?", questionID).First(&params).Error
	if err != nil {
		return nil, err
	}
	return &params, nil
}

// BatchGetParameters implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) BatchGetParameters(ctx context.Context, questionIDs []uint) ([]*models.QuestionParameter, error) {
	var params []*models.QuestionParameter
	err := r.db.WithContext(ctx).Where("question_id IN ?", questionIDs).Find(&params).Error
	return params, err
}

// FindByDifficulty implements repositories.QuestionRepository
func (r *QuestionRepositoryImpl) FindByDifficulty(ctx context.Context, difficulty float64) (*models.Question, error) {
	var question models.Question
	err := r.db.WithContext(ctx).
		Where("difficulty = ?", difficulty).
		First(&question).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}
