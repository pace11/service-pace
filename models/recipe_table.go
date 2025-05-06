package models

import (
	"time"

	"gorm.io/gorm"
)

type RecipeTable struct {
	UUIDModel
	UserID        string         `json:"user_id" gorm:"type:char(36)"`
	Title         string         `json:"title" gorm:"size:255;not null"`
	Content       string         `json:"content" gorm:"type:text;not null"`
	LikesCount    int64          `json:"likes_count" gorm:"column:likes_count"`
	CommentsCount int64          `json:"comments_count" gorm:"column:comments_count"`
	IsLikeByMe    bool           `json:"is_liked_by_me" gorm:"column:is_liked_by_me"`
	IsMine        bool           `json:"is_mine" gorm:"-"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	User UserEmbeddedResponse `json:"user" gorm:"embedded;embeddedPrefix:user__"`
}

func (RecipeTable) TableName() string {
	return "recipes"
}
