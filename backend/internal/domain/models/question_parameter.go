package models

type QuestionParameter struct {
	ID         uint    `gorm:"primarykey"`
	QuestionID uint    `gorm:"not null"`
	Name       string  `gorm:"size:255;not null"`
	MinValue   float64 `gorm:"not null"`
	MaxValue   float64 `gorm:"not null"`
	Step       float64 `gorm:"default:1"`
}
