package models

import (
	"time"

	"gorm.io/gorm"
)

// Question 定义试题模型
type Question struct {
	gorm.Model
	Type            string           `gorm:"type:varchar(20);not null" json:"type"`                       // 题目类型：单选、多选、判断等
	Content         string           `gorm:"type:text;not null" json:"content"`                           // 题目内容
	Answer          string           `gorm:"type:text;not null" json:"answer"`                            // 标准答案
	Analysis        string           `gorm:"type:text" json:"analysis"`                                   // 解析
	SubjectID       uint             `gorm:"not null" json:"subject_id"`                                  // 所属科目
	KnowledgePoints []KnowledgePoint `gorm:"many2many:question_knowledge_points" json:"knowledge_points"` // 关联知识点
	Options         []QuestionOption `gorm:"foreignKey:QuestionID" json:"options"`                        // 选项
	Difficulty      float64          `gorm:"type:decimal(3,2);not null" json:"difficulty"`                // 难度系数
	Score           float64          `gorm:"type:decimal(5,2);not null" json:"score"`                     // 分值
	// IRT参数
	IRTDifficulty     float64 `gorm:"type:decimal(5,2);not null;default:0.5" json:"irt_difficulty"`     // b参数：难度
	IRTDiscrimination float64 `gorm:"type:decimal(5,2);not null;default:1.0" json:"irt_discrimination"` // a参数：区分度
	IRTGuessing       float64 `gorm:"type:decimal(3,2);not null;default:0.0" json:"irt_guessing"`       // c参数：猜测参数
}

// QuestionResponse represents the response for a question
type QuestionResponse struct {
	ID        uint       `json:"id" gorm:"primarykey" swaggertype:"integer"`
	CreatedAt time.Time  `json:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt time.Time  `json:"updated_at" swaggertype:"string" format:"date-time"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time"`
	Question  Question   `json:"question"`
}
