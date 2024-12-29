package models

type QuestionOption struct {
	ID         uint   `gorm:"primarykey"`
	QuestionID uint   `gorm:"not null"`
	Content    string `gorm:"type:text;not null"`
	IsCorrect  bool   `gorm:"default:false"`
}
