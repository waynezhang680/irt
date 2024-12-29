package models

import (
	"gorm.io/gorm"
)

// Subject 定义科目
type Subject struct {
	gorm.Model
	Name            string           `gorm:"uniqueIndex;not null;type:text"`
	Description     string           `gorm:"type:text"`
	Questions       []Question       `gorm:"foreignKey:SubjectID"`
	KnowledgePoints []KnowledgePoint `gorm:"foreignKey:SubjectID"`
	ExamPapers      []ExamPaper      `gorm:"foreignKey:SubjectID"`
	UserAbilities   []UserAbility    `gorm:"foreignKey:SubjectID"`
}
