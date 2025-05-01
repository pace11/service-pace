package repository

import (
	"net/http"
	"service-pace11/config"
	"service-pace11/models"

	"github.com/gin-gonic/gin"
)

type LikeRepository interface {
	Show(id uint) ([]models.LikeResponse, int, string, map[string]string)
	Save(c *gin.Context, id uint) (any, int, string, map[string]string)
	Delete(c *gin.Context, id uint) (any, int, string, map[string]string)
}

type likeRepo struct{}

func NewLikeRepository() LikeRepository {
	return &likeRepo{}
}

func (r *likeRepo) Show(id uint) ([]models.LikeResponse, int, string, map[string]string) {
	var likes []models.Like

	if err := config.DB.Where("recipe_id = ?", id).Order("created_at DESC").Preload("User").Find(&likes).Error; err != nil {
		return nil, http.StatusNotFound, "like", nil
	}

	var response []models.LikeResponse
	for _, l := range likes {
		response = append(response, models.LikeResponse{
			ID:        l.ID,
			RecipeID:  l.RecipeID,
			CreatedAt: l.CreatedAt,
			UpdatedAt: l.UpdatedAt,
			User: models.UserResponse{
				ID:    l.User.ID,
				Name:  l.User.Name,
				Email: l.User.Email,
			},
		})
	}

	return response, http.StatusOK, "like", nil
}

func (r *likeRepo) Save(c *gin.Context, id uint) (any, int, string, map[string]string) {
	var existing models.Like
	userIdRaw, exist := c.Get("user_id")

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	if err := config.DB.Where("user_id = ? AND recipe_id = ?", userIdRaw, id).First(&existing).Error; err != nil {
		return nil, http.StatusNotFound, "like", nil
	}

	if existing.UserID == userIdRaw {
		return nil, http.StatusBadRequest, "you have liked this recipe", nil
	}

	likeRecipe := models.Like{
		UserID:   userIdRaw.(uint),
		RecipeID: id,
	}

	if err := config.DB.Create(&likeRecipe).Error; err != nil {
		return nil, http.StatusInternalServerError, "like", nil
	}

	return nil, http.StatusCreated, "like", nil
}

func (r *likeRepo) Delete(c *gin.Context, id uint) (any, int, string, map[string]string) {
	var existing models.Like
	userIdRaw, exist := c.Get("user_id")

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	if err := config.DB.Where("user_id = ? AND recipe_id = ?", userIdRaw, id).First(&existing).Error; err != nil {
		return nil, http.StatusNotFound, "like", nil
	}

	result := config.DB.Delete(&models.Like{}, existing.ID)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "like", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "like", nil
	}

	return nil, http.StatusOK, "unlike", nil
}
