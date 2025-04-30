package models

type NoteDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RecipeDTO struct {
	UserID  uint   `json:"user_id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
