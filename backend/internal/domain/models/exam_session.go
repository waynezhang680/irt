package models

import (
	"time"

	"gorm.io/gorm"
)

type ExamSession struct {
	gorm.Model
	UserID         uint      `gorm:"not null"`
	StartTime      time.Time `gorm:"not null"`
	EndTime        time.Time
	CurrentAbility float64 // 考生当前能力值估计
	Status         string  `gorm:"default:'in_progress'"`
}
