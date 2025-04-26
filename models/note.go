package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"type:text;not null"`
	Description string         `json:"description" gorm:"type:text;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Note) TableName() string {
	return "notes"
}
