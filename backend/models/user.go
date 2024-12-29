package models

import (
	"time"

	"gorm.io/gorm"
)

// User 定义用户模型
type User struct {
	gorm.Model
	Username     string       `gorm:"uniqueIndex;not null"`
	Password     string       `gorm:"-"`                        // 用于接收密码，但不存储
	PasswordHash string       `gorm:"column:password;not null"` // 存储加密后的密码
	RoleID       uint         `gorm:"not null"`
	RoleType     RoleType     `gorm:"type:varchar(20);not null;default:'student'"`
	Email        string       `gorm:"uniqueIndex;not null"`
	Role         Role         `gorm:"foreignKey:RoleID"`
	ExamRecords  []ExamRecord `gorm:"foreignKey:UserID"`
	Permissions  []string     `json:"permissions,omitempty" gorm:"-"`
	LastLoginAt  *time.Time   `json:"last_login_at,omitempty"`
}

// Role 角色表
type Role struct {
	gorm.Model
	Name        string   `gorm:"unique;not null"`
	Type        RoleType `gorm:"type:varchar(20);not null;default:'student'"`
	Description string
	Users       []User `gorm:"foreignKey:RoleID"`
}
