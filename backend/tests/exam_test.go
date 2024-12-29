package tests

import (
	"irt-exam-system/backend/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamService(t *testing.T) {
	examService := services.NewExamService(db)

	t.Run("CreateExam", func(t *testing.T) {
		exam := &models.Exam{
			Title:     "Test Exam",
			SubjectID: 1,
			Duration:  60,
			PassScore: 60,
		}

		err := examService.CreateExam(exam)
		assert.NoError(t, err)
		assert.NotZero(t, exam.ID)
	})

	t.Run("GetExam", func(t *testing.T) {
		exam, err := examService.GetExam(1)
		assert.NoError(t, err)
		assert.NotNil(t, exam)
	})
}
