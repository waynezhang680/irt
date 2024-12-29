package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ExamRequest struct {
	Title     string  `json:"title" validate:"required,min=2,max=100"`
	SubjectID uint    `json:"subject_id" validate:"required"`
	Duration  int     `json:"duration" validate:"required,min=1,max=180"`
	PassScore float64 `json:"pass_score" validate:"required,min=0"`
}

func ValidateExamRequest(req *ExamRequest) error {
	return validate.Struct(req)
}
