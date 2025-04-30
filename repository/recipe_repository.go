package repository

import (
	"net/http"
	"service-pace11/config"
	"service-pace11/models"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type RecipeRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.Recipe, int, any, int64, int, int)
	Show(c *gin.Context, id uint) (*models.Recipe, int, string, map[string]string)
	// Save(note *models.NoteDTO) (any, int, string, map[string]string)
	// Update(id uint, note *models.NoteDTO) (any, int, string, map[string]string)
	// Delete(id uint) (any, int, string, map[string]string)
}

type recipeRepo struct{}

func NewRecipeRepository() RecipeRepository {
	return &recipeRepo{}
}

func (r *recipeRepo) Index(c *gin.Context, filters map[string]any) ([]models.Recipe, int, any, int64, int, int) {
	userIdRaw, exists := c.Get("user_id")
	var recipes []models.Recipe
	var total int64

	if !exists {
		return nil, http.StatusUnauthorized, "token", 0, 0, 0
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
		Joins("LEFT JOIN (?) as liked_by_me_subquery ON recipes.id = liked_by_me_subquery.recipe_id", likedByMeSubquery)

	query = utils.FilterByParams(query, filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Scan(&recipes)

	for i := range recipes {
		recipes[i].IsMine = recipes[i].UserID == userIdRaw
	}

	return recipes, http.StatusOK, "recipe", total, page, limit
}

func (r *recipeRepo) Show(c *gin.Context, id uint) (*models.Recipe, int, string, map[string]string) {
	userIdRaw, exists := c.Get("user_id")
	var recipe models.Recipe

	if !exists {
		return nil, http.StatusUnauthorized, "token", nil
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

	recipe.IsMine = recipe.UserID == userIdRaw

	return &recipe, http.StatusOK, "recipe", nil
}
