package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUIDModel struct {
	ID string `json:"id" gorm:"type:char(36);primaryKey"`
}

func (b *UUIDModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
