package controllers

import (
	"net/http"
	"service-pace11/models"
	"service-pace11/repository"
	"service-pace11/utils"
	"strconv"

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
	filters := map[string]any{
		"title": c.Query("title"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// GetRecipe
// @Summary Get a single recipe by ID
// @Description Get detail of a recipe by ID
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path int true "Recipe ID"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /recipe/{id} [get]
func (ctl *RecipeController) GetRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Show(c, uint(id))
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

// UpdateRecipe
// @Summary Update a recipe by ID
// @Description Update the title or content of a recipe
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path int true "Recipe ID"
// @Param payload body models.RecipeDTO true "Update recipe payload"
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /recipe/{id} [patch]
func (ctl *RecipeController) UpdateRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var recipe models.RecipeDTO

	if utils.BindAndValidate(c, &recipe) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(c, uint(id), &recipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// DeleteRecipe
// @Summary Delete a recipe by ID
// @Description Delete a recipe from database
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path int true "Recipe ID"
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /recipe/{id} [delete]
func (ctl *RecipeController) DeleteRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
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
func (ctl *RecipeController) ArchiveRecipeIndex(c *gin.Context) {
	data, code, entity, total, page, limit := ctl.Repo.ArchiveIndex(c)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// ArchiveRecipe
// @Summary Create a new recipe save
// @Description Create a new recipe save with title and content
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path int true "Recipe ID"
// @Success 201 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /recipe/save/{id} [post]
func (ctl *RecipeController) ArchiveRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Archive(c, uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
