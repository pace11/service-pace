package repository

import (
	"net/http"
	"service-pace11/config"
	"service-pace11/models"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type CommentRepository interface {
	Show(c *gin.Context, id uint) ([]models.CommentResponse, int, any, int64, int, int)
	Save(c *gin.Context, id uint, comment *models.CommentDTO) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type commentRepo struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepo{}
}

func (r *commentRepo) Show(c *gin.Context, id uint) ([]models.CommentResponse, int, any, int64, int, int) {
	userIDVal, exist := c.Get("user_id")
	var comments []models.Comment
	var total int64

	if !exist {
		return nil, http.StatusUnauthorized, "access", 0, 0, 0
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		return nil, http.StatusInternalServerError, "invalid user_id", 0, 0, 0
	}

	query := config.DB.Model(&models.Comment{}).Where("recipe_id = ?", id).Order("created_at DESC").Preload("User")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&comments)

	var response []models.CommentResponse
	for _, c := range comments {
		response = append(response, models.CommentResponse{
			ID:        c.ID,
			RecipeID:  c.RecipeID,
			Content:   c.Content,
			IsMine:    c.User.ID == userID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			User: models.UserResponse{
				ID:    c.User.ID,
				Name:  c.User.Name,
				Email: c.User.Email,
			},
		})
	}

	return response, http.StatusOK, "comment", total, page, limit
}

func (r *commentRepo) Save(c *gin.Context, id uint, comment *models.CommentDTO) (any, int, string, map[string]string) {
	var existing models.Recipe
	userIDVal, exist := c.Get("user_id")
	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		return nil, http.StatusInternalServerError, "invalid user_id", nil
	}

	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "recipe", nil
	}

	newComment := models.Comment{
		UserID:   userID,
		RecipeID: id,
		Content:  comment.Content,
	}

	if err := config.DB.Create(&newComment).Error; err != nil {
		return nil, http.StatusInternalServerError, "comment", nil
	}

	return nil, http.StatusCreated, "comment", nil
}

func (r *commentRepo) Delete(id uint) (any, int, string, map[string]string) {
	var existing models.Comment

	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "comment", nil
	}

	result := config.DB.Delete(&models.Comment{}, existing.ID)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "comment", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "comment", nil
	}

	return nil, http.StatusOK, "comment", nil
}
