package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	RecipeID  uint      `json:"recipe_id"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Recipe Recipe `json:"recipe" gorm:"foreignKey:RecipeID"`
}

func (Comment) TableName() string {
	return "comments"
}
