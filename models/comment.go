package models

import "time"

type Comment struct {
	UUIDModel
	UserID    string    `json:"user_id" gorm:"type:char(36)"`
	RecipeID  string    `json:"recipe_id" gorm:"type:char(36)"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Recipe Recipe `json:"recipe" gorm:"foreignKey:RecipeID"`
}

func (Comment) TableName() string {
	return "comments"
}
