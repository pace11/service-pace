package models

import "time"

type NotificationType string

const (
	LikeType    NotificationType = "like"
	CommentType NotificationType = "comment"
	SaveType    NotificationType = "save"
)

type Notification struct {
	UUIDModel
	UserID    string           `json:"user_id" gorm:"type:char(36)"`
	Type      NotificationType `json:"type" gorm:"type:enum('like','comment','save');default:'like';not null"`
	Message   string           `json:"message" gorm:"type:text;not null"`
	CreatedAt time.Time        `json:"created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
