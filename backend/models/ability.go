package models

import (
	"time"

	"gorm.io/gorm"
)

// UserAbility 用户能力值表
type UserAbility struct {
	gorm.Model
	ID            uint       `json:"id" gorm:"primarykey" swaggertype:"integer"`
	CreatedAt     time.Time  `json:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt     time.Time  `json:"updated_at" swaggertype:"string" format:"date-time"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time"`
	UserID        uint       `json:"user_id" swaggertype:"integer"`
	SubjectID     uint       `json:"subject_id" swaggertype:"integer"`
	Ability       float64    `json:"ability" swaggertype:"number"`
	StandardError float64    `gorm:"not null;default:0;type:numeric"` // 标准误
	User          User       `gorm:"foreignKey:UserID"`
	Subject       Subject    `gorm:"foreignKey:SubjectID"`
}

// QuestionParameter IRT题目参数表
type QuestionParameter struct {
	gorm.Model
	QuestionID     uint     `gorm:"not null;uniqueIndex"`
	Difficulty     float64  `gorm:"not null;default:0.5;type:numeric"` // b参数：难度
	Discrimination float64  `gorm:"not null;default:1.0;type:numeric"` // a参数：区分度
	Guessing       float64  `gorm:"not null;default:0.0;type:numeric"` // c参数：猜测参数
	Question       Question `gorm:"foreignKey:QuestionID"`
}

// AbilityEstimation 能力值估计历史记录
type AbilityEstimation struct {
	gorm.Model
	UserID        uint       `gorm:"not null;index"`
	SubjectID     uint       `gorm:"not null;index"`
	ExamRecordID  uint       `gorm:"not null;index"`
	Ability       float64    `gorm:"not null;type:numeric"` // 估计的能力值
	StandardError float64    `gorm:"not null;type:numeric"` // 标准误
	Method        string     `gorm:"not null;type:text"`    // 估计方法（如：MLE, EAP, MAP等）
	User          User       `gorm:"foreignKey:UserID"`
	Subject       Subject    `gorm:"foreignKey:SubjectID"`
	ExamRecord    ExamRecord `gorm:"foreignKey:ExamRecordID"`
}
