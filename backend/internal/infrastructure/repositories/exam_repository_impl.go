package repositories

import (
	"context"
	"errors"

	"irt-exam-system/backend/internal/domain/models"
	"irt-exam-system/backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type examRepository struct {
	db *gorm.DB
}

// NewExamRepository 创建考试仓储实例
func NewExamRepository(db *gorm.DB) repositories.ExamRepository {
	return &examRepository{db: db}
}

// 试卷相关实现
func (r *examRepository) CreatePaper(ctx context.Context, paper *models.ExamPaper) error {
	return r.db.WithContext(ctx).Create(paper).Error
}

func (r *examRepository) UpdatePaper(ctx context.Context, paper *models.ExamPaper) error {
	return r.db.WithContext(ctx).Save(paper).Error
}

func (r *examRepository) DeletePaper(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ExamPaper{}, id).Error
}

func (r *examRepository) FindPaperByID(ctx context.Context, id uint) (*models.ExamPaper, error) {
	var paper models.ExamPaper
	err := r.db.WithContext(ctx).Preload("Questions").First(&paper, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &paper, nil
}

func (r *examRepository) ListPapers(ctx context.Context, offset, limit int) ([]*models.ExamPaper, int64, error) {
	var papers []*models.ExamPaper
	var total int64

	err := r.db.WithContext(ctx).Model(&models.ExamPaper{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Preload("Questions").Offset(offset).Limit(limit).Find(&papers).Error
	if err != nil {
		return nil, 0, err
	}

	return papers, total, nil
}

func (r *examRepository) ListPapersBySubject(ctx context.Context, subjectID uint, offset, limit int) ([]*models.ExamPaper, int64, error) {
	var papers []*models.ExamPaper
	var total int64

	err := r.db.WithContext(ctx).Model(&models.ExamPaper{}).Where("subject_id = ?", subjectID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("subject_id = ?", subjectID).
		Preload("Questions").Offset(offset).Limit(limit).Find(&papers).Error
	if err != nil {
		return nil, 0, err
	}

	return papers, total, nil
}

func (r *examRepository) ListPapersByStatus(ctx context.Context, status string, offset, limit int) ([]*models.ExamPaper, int64, error) {
	var papers []*models.ExamPaper
	var total int64

	err := r.db.WithContext(ctx).Model(&models.ExamPaper{}).Where("status = ?", status).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("status = ?", status).
		Preload("Questions").Offset(offset).Limit(limit).Find(&papers).Error
	if err != nil {
		return nil, 0, err
	}

	return papers, total, nil
}

// 考试记录相关实现
func (r *examRepository) CreateRecord(ctx context.Context, record *models.ExamSession) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *examRepository) UpdateRecord(ctx context.Context, record *models.ExamSession) error {
	return r.db.WithContext(ctx).Save(record).Error
}

func (r *examRepository) FindRecordByID(ctx context.Context, id uint) (*models.ExamSession, error) {
	var record models.ExamSession
	err := r.db.WithContext(ctx).Preload("Responses").First(&record, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (r *examRepository) ListUserRecords(ctx context.Context, userID uint, offset, limit int) ([]*models.ExamSession, int64, error) {
	var records []*models.ExamSession
	var total int64

	err := r.db.WithContext(ctx).Model(&models.ExamSession{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("user_id = ?", userID).
		Preload("Responses").Offset(offset).Limit(limit).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (r *examRepository) ListPaperRecords(ctx context.Context, paperID uint, offset, limit int) ([]*models.ExamSession, int64, error) {
	var records []*models.ExamSession
	var total int64

	err := r.db.WithContext(ctx).Model(&models.ExamSession{}).Where("exam_paper_id = ?", paperID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("exam_paper_id = ?", paperID).
		Preload("Responses").Offset(offset).Limit(limit).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// 答题记录相关实现
func (r *examRepository) CreateResponse(ctx context.Context, response *models.QuestionResponse) error {
	return r.db.WithContext(ctx).Create(response).Error
}

func (r *examRepository) BatchCreateResponses(ctx context.Context, responses []*models.QuestionResponse) error {
	return r.db.WithContext(ctx).Create(&responses).Error
}

func (r *examRepository) UpdateResponse(ctx context.Context, response *models.QuestionResponse) error {
	return r.db.WithContext(ctx).Save(response).Error
}

func (r *examRepository) ListRecordResponses(ctx context.Context, recordID uint) ([]*models.QuestionResponse, error) {
	var responses []*models.QuestionResponse
	err := r.db.WithContext(ctx).Where("exam_record_id = ?", recordID).Find(&responses).Error
	return responses, err
}

func (r *examRepository) GetResponseStats(ctx context.Context, questionID uint) (total int64, correct int64, error error) {
	err := r.db.WithContext(ctx).Model(&models.ExamResponse{}).
		Where("question_id = ?", questionID).Count(&total).Error
	if err != nil {
		return 0, 0, err
	}

	err = r.db.WithContext(ctx).Model(&models.ExamResponse{}).
		Where("question_id = ? AND is_correct = ?", questionID, true).Count(&correct).Error
	if err != nil {
		return 0, 0, err
	}

	return total, correct, nil
}
