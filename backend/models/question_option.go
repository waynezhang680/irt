package models

import (
	"gorm.io/gorm"
)

// QuestionOption 定义题目选项
type QuestionOption struct {
	gorm.Model
	QuestionID uint     `gorm:"not null;index"`
	Content    string   `gorm:"not null;type:text"`
	IsCorrect  bool     `gorm:"not null"`
	Label      string   `gorm:"not null;type:text"`   // 选项标签，如 A, B, C, D
	Order      int64    `gorm:"not null;type:bigint"` // 选项顺序
	Question   Question `gorm:"foreignKey:QuestionID"`
}
