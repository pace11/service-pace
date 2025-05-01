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

func (ctl *LikeController) GetLikeByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Show(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *LikeController) LikeByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Save(c, uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *LikeController) UnlikeByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Delete(c, uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
