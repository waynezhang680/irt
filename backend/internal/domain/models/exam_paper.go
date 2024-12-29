package models

import "time"

type ExamPaper struct {
	ID          uint    `gorm:"primarykey"`
	Title       string  `gorm:"size:255;not null"`
	Description string  `gorm:"type:text"`
	Duration    int     `gorm:"not null"` // in minutes
	TotalScore  float64 `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
