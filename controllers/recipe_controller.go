package controllers

import (
	"net/http"
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

func (ctl *RecipeController) GetRecipes(c *gin.Context) {
	filters := map[string]any{
		"title": c.Query("title"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *RecipeController) GetRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Show(c, uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
