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
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CommentDTO struct {
	Content string `json:"content" binding:"required"`
}

type RegisterDTO struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateDTO struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
}
