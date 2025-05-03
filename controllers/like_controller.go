package controllers

import (
	"net/http"
	"service-pace11/repository"
	"service-pace11/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	Repo repository.LikeRepository
}

func NewLikeController(repo repository.LikeRepository) *LikeController {
	return &LikeController{Repo: repo}
}

// GetLikesByRecipe
// @Summary Get list of likes by recipe
// @Description Retrieve all likes with pagination
// @Tags Likes
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /likes/recipe/{id} [get]
func (ctl *LikeController) GetLikesByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, total, page, limit := ctl.Repo.Show(c, uint(id))
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// LikeByRecipe
// @Summary Create a like by recipe
// @Description Create a new like by recipe
// @Tags Likes
// @Accept json
// @Produce json
// @Success 201 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /like/recipe/{id} [post]
func (ctl *LikeController) LikeByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Save(c, uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// UnlikeByRecipe
// @Summary Create an unlike by recipe
// @Description Create an unlike by recipe
// @Tags Likes
// @Accept json
// @Produce json
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /unlike/recipe/{id} [post]
func (ctl *LikeController) UnlikeByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Delete(c, uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
