package models

import (
	"gorm.io/gorm"
)

// KnowledgePoint 定义知识点
type KnowledgePoint struct {
	gorm.Model
	SubjectID   uint                     `gorm:"not null;index"`
	Name        string                   `gorm:"not null;type:text"`
	Description string                   `gorm:"type:text"`
	ParentID    *uint                    `gorm:"index"` // 父知识点ID，允许为空
	Subject     Subject                  `gorm:"foreignKey:SubjectID"`
	Parent      *KnowledgePoint          `gorm:"foreignKey:ParentID"`
	Children    []KnowledgePoint         `gorm:"foreignKey:ParentID"`
	Questions   []QuestionKnowledgePoint `gorm:"foreignKey:KnowledgePointID"`
}

// QuestionKnowledgePoint 定义题目和知识点的多对多关系
type QuestionKnowledgePoint struct {
	gorm.Model
	QuestionID       uint           `gorm:"not null;index"`
	KnowledgePointID uint           `gorm:"not null;index"`
	Question         Question       `gorm:"foreignKey:QuestionID"`
	KnowledgePoint   KnowledgePoint `gorm:"foreignKey:KnowledgePointID"`
}
