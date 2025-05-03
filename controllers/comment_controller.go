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

// GetCommentsByRecipe
// @Summary Get list of comments by recipe
// @Description Retrieve all comments with pagination
// @Tags Comments
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /comments/recipe/{id} [get]
func (ctl *CommentController) GetCommentsByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	id, err := strconv.Atoi(idRecipe)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, total, page, limit := ctl.Repo.Show(c, uint(id))
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// CommentByRecipe
// @Summary Create a comment by recipe
// @Description Create a new comment by recipe
// @Tags Comments
// @Accept json
// @Produce json
// @Success 201 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /comment/recipe/{id} [post]
func (ctl *CommentController) CommentByRecipe(c *gin.Context) {
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

// DeleteComment
// @Summary Delete a comment by ID
// @Description Delete a comment from database
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /comment/{id} [delete]
func (ctl *CommentController) DeleteComment(c *gin.Context) {
	idComment := c.Param("id")
	id, err := strconv.Atoi(idComment)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
	}

	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
