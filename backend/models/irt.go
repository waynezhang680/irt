package models

import (
	"gorm.io/gorm"
)

type QuestionDifficulty struct {
	gorm.Model
	QuestionID     uint    `gorm:"not null"`
	Difficulty     float64 `gorm:"not null"` // 题目难度参数
	Discrimination float64 `gorm:"not null"` // 题目区分度参数
	Guessing       float64 // 猜测参数
}
