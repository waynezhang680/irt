package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content        string  `gorm:"type:text;not null"`
	OptionA        string  `gorm:"not null"`
	OptionB        string  `gorm:"not null"`
	OptionC        string  `gorm:"not null"`
	OptionD        string  `gorm:"not null"`
	Answer         string  `gorm:"not null"`
	Difficulty     float64 `gorm:"not null"` // IRT难度参数
	Discrimination float64 `gorm:"not null"` // IRT区分度参数
	GuessParameter float64 `gorm:"not null"` // IRT猜测参数
	Score          float64 `gorm:"not null"` // Add this field
}

type QuestionKnowledgePoint struct {
	ID               uint `gorm:"primaryKey"`
	QuestionID       uint
	KnowledgePointID uint
}

func (QuestionKnowledgePoint) TableName() string {
	return "question_knowledge_points"
}
