package models

import "time"

type ExamResponse struct {
	ID         uint    `gorm:"primarykey"`
	RecordID   uint    `gorm:"not null"`
	QuestionID uint    `gorm:"not null"`
	Answer     string  `gorm:"type:text"`
	Score      float64 `gorm:"default:0"`
	TimeSpent  int64   `gorm:"default:0"` // in seconds
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
