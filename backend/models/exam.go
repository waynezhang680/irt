package models

import (
	"time"

	"gorm.io/gorm"
)

// ExamPaper 定义考试试卷
type ExamPaper struct {
	gorm.Model
	Title       string              `gorm:"not null;type:text"`
	SubjectID   uint                `gorm:"not null;index"`
	Description string              `gorm:"type:text"`
	TimeLimit   int64               `gorm:"not null;type:bigint"`
	TotalScore  float64             `gorm:"not null;type:numeric"`
	PassScore   float64             `gorm:"not null;type:numeric"`
	Status      string              `gorm:"not null;default:'draft';type:text"`
	StartTime   time.Time           `gorm:"type:timestamptz"`
	EndTime     time.Time           `gorm:"type:timestamptz"`
	Subject     Subject             `gorm:"foreignKey:SubjectID"`
	Questions   []ExamPaperQuestion `gorm:"foreignKey:ExamPaperID"`
	Records     []ExamRecord        `gorm:"foreignKey:ExamPaperID"`
}

// ExamPaperQuestion 定义试卷题目关联
type ExamPaperQuestion struct {
	gorm.Model
	ExamPaperID uint      `gorm:"not null;index"`
	QuestionID  uint      `gorm:"not null;index"`
	Score       float64   `gorm:"not null;type:numeric"`
	Order       int64     `gorm:"not null;type:bigint"`
	ExamPaper   ExamPaper `gorm:"foreignKey:ExamPaperID"`
	Question    Question  `gorm:"foreignKey:QuestionID"`
}

// ExamRecord 定义考试记录
type ExamRecord struct {
	gorm.Model
	UserID      uint           `gorm:"not null;index"`
	ExamPaperID uint           `gorm:"not null;index"`
	StartTime   time.Time      `gorm:"not null;type:timestamptz"`
	EndTime     time.Time      `gorm:"type:timestamptz"`
	Score       float64        `gorm:"type:numeric"`
	Status      string         `gorm:"not null;default:'in_progress';type:text"`
	User        User           `gorm:"foreignKey:UserID"`
	ExamPaper   ExamPaper      `gorm:"foreignKey:ExamPaperID"`
	Responses   []ExamResponse `gorm:"foreignKey:ExamRecordID"`
}

// ExamResponse 定义答题记录
type ExamResponse struct {
	gorm.Model
	ExamRecordID uint       `gorm:"not null;index"`
	QuestionID   uint       `gorm:"not null;index"`
	UserAnswer   string     `gorm:"not null;type:text"`
	IsCorrect    bool       `gorm:"not null"`
	ResponseTime int64      `gorm:"not null;type:bigint"` // 答题用时（秒）
	Score        float64    `gorm:"not null;type:numeric"`
	ExamRecord   ExamRecord `gorm:"foreignKey:ExamRecordID"`
	Question     Question   `gorm:"foreignKey:QuestionID"`
}

// TotalQuestions 获取总题目数
func (r *ExamRecord) TotalQuestions() int {
	return len(r.Responses)
}

// CorrectCount 获取正确题目数
func (r *ExamRecord) CorrectCount() int {
	count := 0
	for _, response := range r.Responses {
		if response.IsCorrect {
			count++
		}
	}
	return count
}

// TimeSpent 获取考试用时（秒）
func (r *ExamRecord) TimeSpent() int64 {
	if r.EndTime.IsZero() {
		return 0
	}
	return int64(r.EndTime.Sub(r.StartTime).Seconds())
}
