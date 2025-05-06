package repository

import (
	"errors"
	"fmt"
	"net/http"
	"service-pace11/config"
	"service-pace11/models"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LikeRepository interface {
	Show(c *gin.Context, id string) ([]models.LikeResponse, int, any, int64, int, int)
	Save(c *gin.Context, id string) (any, int, string, map[string]string)
	Delete(c *gin.Context, id string) (any, int, string, map[string]string)
}

type likeRepo struct {
	notificationRepo NotificationRepository
}

func NewLikeRepository(notificationRepo NotificationRepository) LikeRepository {
	return &likeRepo{
		notificationRepo: notificationRepo,
	}
}

func (r *likeRepo) Show(c *gin.Context, id string) ([]models.LikeResponse, int, any, int64, int, int) {
	var likes []models.Like
	var total int64

	query := config.DB.Model(&models.Like{}).Where("recipe_id = ?", id).Order("created_at DESC").Preload("User")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&likes)

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

	return response, http.StatusOK, "like", total, page, limit
}

func (r *likeRepo) Save(c *gin.Context, id string) (any, int, string, map[string]string) {
	var detailRecipe models.Recipe
	userIDVal, exist := c.Get("user_id")
	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	userID, ok := userIDVal.(string)
	if !ok {
		return nil, http.StatusInternalServerError, "invalid user_id", nil
	}

	config.DB.Where("id = ? ", id).Preload("User").Find(&detailRecipe)

	var existing models.Like
	err := config.DB.Where("user_id = ? AND recipe_id = ?", userID, id).First(&existing).Error
	if err == nil {
		return nil, http.StatusBadRequest, "you have liked this recipe", nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, "like", nil
	}

	newLike := models.Like{
		UserID:   userID,
		RecipeID: id,
	}

	if err := config.DB.Create(&newLike).Error; err != nil {
		return nil, http.StatusInternalServerError, "like", nil
	}

	created := r.notificationRepo.Save(c, detailRecipe.UserID, userID, string(models.LikeType), detailRecipe.Title)
	if !created {
		fmt.Print("Notification not created")
	}

	return nil, http.StatusCreated, "like", nil
}

func (r *likeRepo) Delete(c *gin.Context, id string) (any, int, string, map[string]string) {
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
