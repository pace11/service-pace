package models

import "time"

type NoteResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserDetailResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserEmbeddedResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RecipeResponse struct {
	ID            string               `json:"id"`
	Title         string               `json:"title"`
	Content       string               `json:"content"`
	LikesCount    int64                `json:"likes_count" gorm:"column:likes_count"`
	CommentsCount int64                `json:"comments_count" gorm:"column:comments_count"`
	IsLikeByMe    bool                 `json:"is_liked_by_me" gorm:"column:is_liked_by_me"`
	IsSaveByMe    bool                 `json:"is_saved_by_me" gorm:"column:is_saved_by_me"`
	IsMine        bool                 `json:"is_mine" gorm:"-"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	User          UserEmbeddedResponse `json:"user" gorm:"embedded;embeddedPrefix:user__"`
}

type LikeResponse struct {
	ID        string       `json:"id"`
	RecipeID  string       `json:"recipe_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserResponse `json:"user"`
}

type CommentResponse struct {
	ID        string       `json:"id"`
	RecipeID  string       `json:"recipe_id"`
	Content   string       `json:"content"`
	IsMine    bool         `json:"is_mine" gorm:"-"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserResponse `json:"user"`
}
