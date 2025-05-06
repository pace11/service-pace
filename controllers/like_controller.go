package controllers

import (
	"service-pace11/repository"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	Repo repository.LikeRepository
}

func NewLikeController(repo repository.LikeRepository) *LikeController {
	return &LikeController{Repo: repo}
}

// GetLikesByRecipeID
// @Summary Get list of likes by recipe ID
// @Description Retrieve all likes with pagination
// @Tags Likes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /likes/recipe/{id} [get]
func (ctl *LikeController) GetLikesByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")

	data, code, entity, total, page, limit := ctl.Repo.Show(c, idRecipe)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// LikeByRecipeID
// @Summary Create a like by recipe ID
// @Description Create a new like by recipe ID
// @Tags Likes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Success 201 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /like/recipe/{id} [post]
func (ctl *LikeController) LikeByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")

	data, code, entity, errors := ctl.Repo.Save(c, idRecipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// UnlikeByRecipeID
// @Summary Create an unlike by recipe ID
// @Description Create an unlike by recipe ID
// @Tags Likes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /unlike/recipe/{id} [post]
func (ctl *LikeController) UnlikeByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")

	data, code, entity, errors := ctl.Repo.Delete(c, idRecipe)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
