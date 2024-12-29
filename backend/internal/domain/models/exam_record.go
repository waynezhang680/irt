package models

import "time"

type ExamRecord struct {
	ID         uint    `gorm:"primarykey"`
	UserID     uint    `gorm:"not null"`
	ExamID     uint    `gorm:"not null"`
	Score      float64 `gorm:"default:0"`
	Status     string  `gorm:"size:20;not null"` // e.g., "in_progress", "completed"
	StartTime  time.Time
	EndTime    *time.Time
	TotalTime  int64 `gorm:"default:0"` // in seconds
	AutoSubmit bool  `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
