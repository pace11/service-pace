package models

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	UUIDModel
	UserID    string         `json:"user_id" gorm:"type:char(36)"`
	Title     string         `json:"title" gorm:"size:255;not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}

func (Recipe) TableName() string {
	return "recipes"
}
