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

type RecipeRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.RecipeResponse, int, any, int64, int, int)
	Show(c *gin.Context, id string) (*models.RecipeResponse, int, string, map[string]string)
	Save(c *gin.Context, recipe *models.RecipeDTO) (any, int, string, map[string]string)
	Update(c *gin.Context, id string, recipe *models.RecipeDTO) (any, int, string, map[string]string)
	Delete(id string) (any, int, string, map[string]string)
	SavedIndex(c *gin.Context) ([]models.SavedRecipe, int, any, int64, int, int)
	Saved(c *gin.Context, id string) (any, int, string, map[string]string)
}

type recipeRepo struct {
	notificationRepo NotificationRepository
}

func NewRecipeRepository(notificationRepo NotificationRepository) RecipeRepository {
	return &recipeRepo{
		notificationRepo: notificationRepo,
	}
}

func (r *recipeRepo) Index(c *gin.Context, filters map[string]any) ([]models.RecipeResponse, int, any, int64, int, int) {
	userIdRaw, exist := c.Get("user_id")
	var recipes []models.RecipeTable
	var total int64

	if !exist {
		return nil, http.StatusUnauthorized, "access", 0, 0, 0
	}

	likesSubquery := config.DB.
		Table("likes").
		Select("recipe_id, COUNT(*) as likes_count").
		Group("recipe_id")

	commentsSubQuery := config.DB.
		Table("comments").
		Select("recipe_id, COUNT(*) as comments_count").
		Group("recipe_id")

	likedByMeSubquery := config.DB.
		Table("likes").
		Select("recipe_id, true as is_liked_by_me").
		Where("user_id = ?", userIdRaw)

	query := config.DB.Model(&models.RecipeTable{}).
		Select(`
			recipes.*,
			users.id as user__id,
			users.name as user__name,
			users.email as user__email,
			IFNULL(likes_subquery.likes_count, 0) as likes_count,
			IFNULL(comments_subquery.comments_count, 0) as comments_count,
			IFNULL(liked_by_me_subquery.is_liked_by_me, false) as is_liked_by_me
		`).
		Joins("LEFT JOIN users ON users.id = recipes.user_id").
		Joins("LEFT JOIN (?) as likes_subquery ON recipes.id = likes_subquery.recipe_id", likesSubquery).
		Joins("LEFT JOIN (?) as comments_subquery ON recipes.id = comments_subquery.recipe_id", commentsSubQuery).
		Joins("LEFT JOIN (?) as liked_by_me_subquery ON recipes.id = liked_by_me_subquery.recipe_id", likedByMeSubquery).
		Order("created_at DESC")

	query = utils.FilterByParams(query, filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Scan(&recipes)

	var response []models.RecipeResponse
	for _, r := range recipes {
		response = append(response, models.RecipeResponse{
			ID:            r.ID,
			Title:         r.Title,
			Content:       r.Content,
			LikesCount:    r.LikesCount,
			CommentsCount: r.CommentsCount,
			IsLikeByMe:    r.IsLikeByMe,
			IsMine:        r.UserID == userIdRaw,
			CreatedAt:     r.CreatedAt,
			UpdatedAt:     r.UpdatedAt,
			User: models.UserEmbeddedResponse{
				ID:    r.User.ID,
				Name:  r.User.Name,
				Email: r.User.Email,
			},
		})
	}

	return response, http.StatusOK, "recipe", total, page, limit
}

func (r *recipeRepo) Show(c *gin.Context, id string) (*models.RecipeResponse, int, string, map[string]string) {
	userIdRaw, exist := c.Get("user_id")
	var recipe models.RecipeTable

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	likesSubquery := config.DB.
		Table("likes").
		Select("recipe_id, COUNT(*) as likes_count").
		Group("recipe_id")

	commentsSubQuery := config.DB.
		Table("comments").
		Select("recipe_id, COUNT(*) as comments_count").
		Group("recipe_id")

	likedByMeSubquery := config.DB.
		Table("likes").
		Select("recipe_id, true as is_liked_by_me").
		Where("user_id = ?", userIdRaw)

	query := config.DB.Model(&models.Recipe{}).
		Select(`
			recipes.*,
			users.id as user__id,
			users.name as user__name,
			users.email as user__email,
			IFNULL(likes_subquery.likes_count, 0) as likes_count,
			IFNULL(comments_subquery.comments_count, 0) as comments_count,
			IFNULL(liked_by_me_subquery.is_liked_by_me, false) as is_liked_by_me
		`).
		Joins("LEFT JOIN users ON users.id = recipes.user_id").
		Joins("LEFT JOIN (?) as likes_subquery ON recipes.id = likes_subquery.recipe_id", likesSubquery).
		Joins("LEFT JOIN (?) as comments_subquery ON recipes.id = comments_subquery.recipe_id", commentsSubQuery).
		Joins("LEFT JOIN (?) as liked_by_me_subquery ON recipes.id = liked_by_me_subquery.recipe_id", likedByMeSubquery).
		Where("recipes.id = ?", id)

	if err := query.Scan(&recipe).Error; err != nil {
		return nil, http.StatusNotFound, "recipe", nil
	}

	response := &models.RecipeResponse{
		ID:            recipe.ID,
		Title:         recipe.Title,
		Content:       recipe.Content,
		LikesCount:    recipe.LikesCount,
		CommentsCount: recipe.CommentsCount,
		IsLikeByMe:    recipe.IsLikeByMe,
		IsMine:        recipe.UserID == userIdRaw,
		User: models.UserEmbeddedResponse{
			ID:    recipe.User.ID,
			Name:  recipe.User.Name,
			Email: recipe.User.Email,
		},
		CreatedAt: recipe.CreatedAt,
		UpdatedAt: recipe.UpdatedAt,
	}

	return response, http.StatusOK, "recipe", nil
}

func (r *recipeRepo) Save(c *gin.Context, recipe *models.RecipeDTO) (any, int, string, map[string]string) {
	userIdRaw, exist := c.Get("user_id")

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	userID, ok := userIdRaw.(string)
	if !ok {
		return nil, http.StatusInternalServerError, "invalid user_id", nil
	}

	recipeCreate := models.Recipe{
		UserID:  userID,
		Title:   recipe.Title,
		Content: recipe.Content,
	}

	if err := config.DB.Create(&recipeCreate).Error; err != nil {
		return nil, http.StatusInternalServerError, "recipe", nil
	}

	return recipe, http.StatusCreated, "recipe", nil
}

func (r *recipeRepo) Update(c *gin.Context, id string, recipe *models.RecipeDTO) (any, int, string, map[string]string) {
	userIdRaw, exist := c.Get("user_id")
	var existing models.Recipe

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	if err := config.DB.Where("id = ? AND user_id = ?", id, userIdRaw).First(&existing).Error; err != nil {
		return nil, http.StatusNotFound, "recipe", nil
	}

	if err := config.DB.Model(&existing).Updates(map[string]any{
		"title":   recipe.Title,
		"content": recipe.Content,
	}).Error; err != nil {
		return nil, http.StatusInternalServerError, "recipe", nil
	}

	return recipe, http.StatusOK, "recipe", nil
}

func (r *recipeRepo) Delete(id string) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.Recipe{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "recipe", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "recipe", nil
	}

	return nil, http.StatusOK, "recipe", nil
}

func (r *recipeRepo) SavedIndex(c *gin.Context) ([]models.SavedRecipe, int, any, int64, int, int) {
	userIdRaw, exist := c.Get("user_id")
	var saveRecipes []models.SavedRecipe
	var total int64

	if !exist {
		return nil, http.StatusUnauthorized, "access", 0, 0, 0
	}

	query := config.DB.Model(&models.SavedRecipe{}).Where("user_id = ?", userIdRaw).Preload("User").Preload("Recipe").Order("created_at DESC")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&saveRecipes)

	var response []models.SavedRecipe
	for _, s := range saveRecipes {
		response = append(response, models.SavedRecipe{
			UUIDModel: models.UUIDModel{ID: s.ID},
			UserID:    s.UserID,
			RecipeID:  s.RecipeID,
			User:      s.User,
			Recipe:    s.Recipe,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		})
	}

	return response, http.StatusOK, "save recipes", total, page, limit
}

func (r *recipeRepo) Saved(c *gin.Context, id string) (any, int, string, map[string]string) {
	var detailRecipe models.Recipe
	var saveRecipe models.SavedRecipe
	userIdRaw, exist := c.Get("user_id")

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	userID, ok := userIdRaw.(string)
	if !ok {
		return nil, http.StatusInternalServerError, "invalid user_id", nil
	}

	config.DB.Where("id = ? ", id).Preload("User").Find(&detailRecipe)

	err := config.DB.Where("user_id = ? AND recipe_id = ?", userIdRaw, id).First(&saveRecipe).Error
	if err == nil {
		return nil, http.StatusBadRequest, "you have saved this recipe", nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, "save recipe", nil
	}

	newSaveRecipe := models.SavedRecipe{
		UserID:   userID,
		RecipeID: id,
	}

	if err := config.DB.Create(&newSaveRecipe).Error; err != nil {
		return nil, http.StatusInternalServerError, "save recipe", nil
	}

	created := r.notificationRepo.Save(c, detailRecipe.UserID, userID, string(models.SaveType), detailRecipe.Title)
	if !created {
		fmt.Print("Notification not created")
	}

	return nil, http.StatusOK, "save recipe", nil
}
