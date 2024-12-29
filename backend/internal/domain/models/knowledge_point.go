package models

type KnowledgePoint struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	ParentID    *uint  `gorm:"default:null"` // For hierarchical structure
}
