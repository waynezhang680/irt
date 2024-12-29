package repositories

import (
	"irt-exam-system/backend/internal/domain/models"
	"gorm.io/gorm"
)

type ExamSessionRepositoryImpl struct {
	db *gorm.DB
}

func NewExamSessionRepository(db *gorm.DB) *ExamSessionRepositoryImpl {
	return &ExamSessionRepositoryImpl{db: db}
}

func (r *ExamSessionRepositoryImpl) Create(session *models.ExamSession) error {
	return r.db.Create(session).Error
}

func (r *ExamSessionRepositoryImpl) FindByID(id uint) (*models.ExamSession, error) {
	var session models.ExamSession
	if err := r.db.First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *ExamSessionRepositoryImpl) SaveResponse(response *models.QuestionResponse) error {
	return r.db.Create(response).Error
}

func (r *ExamSessionRepositoryImpl) UpdateAbility(sessionID uint, ability float64) error {
	return r.db.Model(&models.ExamSession{}).
		Where("id = ?", sessionID).
		Update("current_ability", ability).Error
}
