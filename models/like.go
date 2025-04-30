package models

import "time"

type Like struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	RecipeID  uint      `json:"recipe_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Recipe Recipe `json:"recipe" gorm:"foreignKey:RecipeID"`
}

func (Like) TableName() string {
	return "likes"
}
