package models

import "time"

type SavedRecipe struct {
	UUIDModel
	UserID    string    `json:"user_id" gorm:"type:char(36)"`
	RecipeID  string    `json:"recipe_id" gorm:"type:char(36)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Recipe Recipe `json:"recipe" gorm:"foreignKey:RecipeID"`
}

func (SavedRecipe) TableName() string {
	return "save_recipes"
}
