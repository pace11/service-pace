package controllers

import (
	"service-pace11/models"
	"service-pace11/repository"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type RecipeController struct {
	Repo repository.RecipeRepository
}

func NewRecipeController(repo repository.RecipeRepository) *RecipeController {
	return &RecipeController{Repo: repo}
}

// GetRecipes
// @Summary Get list of recipes
// @Description Retrieve all recipes with pagination
// @Tags Recipes
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /recipes [get]
func (ctl *RecipeController) GetRecipes(c *gin.Context) {
	userIdRaw, exist := c.Get("user_id")

	filters := map[string]any{
		"title": c.Query("title"),
	}

	if !exist {
		return
	}

	if c.Query("type") == "me" {
		filters["user_id"] = userIdRaw
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// GetRecipeByID
// @Summary Get a single recipe by ID
// @Description Get detail of a recipe by ID
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /recipe/{id} [get]
func (ctl *RecipeController) GetRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	data, code, entity, errors := ctl.Repo.Show(c, idRecipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// CreateRecipe
// @Summary Create a new recipe
// @Description Create a new recipe with title and content
// @Tags Recipes
// @Accept json
// @Produce json
// @Param payload body models.RecipeDTO true "Create recipe payload"
// @Success 201 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /recipe [post]
func (ctl *RecipeController) CreateRecipe(c *gin.Context) {
	var recipe models.RecipeDTO

	if utils.BindAndValidate(c, &recipe) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(c, &recipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// UpdateRecipeByID
// @Summary Update a recipe by ID
// @Description Update the title or content of a recipe
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Param payload body models.RecipeDTO true "Update recipe payload"
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /recipe/{id} [patch]
func (ctl *RecipeController) UpdateRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	var recipe models.RecipeDTO

	if utils.BindAndValidate(c, &recipe) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(c, idRecipe, &recipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// DeleteRecipeByID
// @Summary Delete a recipe by ID
// @Description Delete a recipe from database
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /recipe/{id} [delete]
func (ctl *RecipeController) DeleteRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	data, code, entity, errors := ctl.Repo.Delete(idRecipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// GetRecipeSaves
// @Summary Get list of recipes save
// @Description Retrieve all recipes save with pagination
// @Tags Recipes
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /recipe/saves [get]
func (ctl *RecipeController) SavedRecipeIndex(c *gin.Context) {
	data, code, entity, total, page, limit := ctl.Repo.SavedIndex(c)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// ArchiveRecipe
// @Summary Create a new recipe save
// @Description Create a new recipe save with title and content
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Success 201 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /recipe/save/{id} [post]
func (ctl *RecipeController) SavedRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	data, code, entity, errors := ctl.Repo.Saved(c, idRecipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
