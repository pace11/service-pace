package controllers

import (
	"net/http"
	"service-pace11/models"
	"service-pace11/repository"
	"service-pace11/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	Repo repository.CommentRepository
}

func NewCommentController(repo repository.CommentRepository) *CommentController {
	return &CommentController{Repo: repo}
}

func (ctl *CommentController) GetCommentsByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, total, page, limit := ctl.Repo.Show(c, uint(id))
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *CommentController) CreateComment(c *gin.Context) {
	var comment models.CommentDTO
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if utils.BindAndValidate(c, &comment) != nil {
		return
	}

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Save(c, uint(id), &comment)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *CommentController) DeleteComment(c *gin.Context) {
	idComment := c.Param("id")
	id, err := strconv.Atoi(idComment)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
