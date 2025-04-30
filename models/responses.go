package models

import "time"

type NoteResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserEmbeddedResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RecipeResponse struct {
	ID         uint                 `json:"id"`
	Title      string               `json:"title"`
	Content    string               `json:"content"`
	LikesCount int64                `json:"likes_count"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	User       UserEmbeddedResponse `json:"user" gorm:"embedded;embeddedPrefix:user__"`
}
