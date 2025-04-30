package models

type NoteDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
