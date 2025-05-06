package controllers

import (
	"service-pace11/models"
	"service-pace11/repository"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	Repo repository.CommentRepository
}

func NewCommentController(repo repository.CommentRepository) *CommentController {
	return &CommentController{Repo: repo}
}

// GetCommentsByRecipeID
// @Summary Get list of comments by recipe ID
// @Description Retrieve all comments with pagination
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /comments/recipe/{id} [get]
func (ctl *CommentController) GetCommentsByRecipe(c *gin.Context) {
	idRecipe := c.Param("id")
	data, code, entity, total, page, limit := ctl.Repo.Show(c, idRecipe)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

// CommentByRecipeID
// @Summary Create a comment by recipe ID
// @Description Create a new comment by recipe ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID (UUID)" format(uuid)
// @Param payload body models.CommentDTO true "Comment payload"
// @Success 201 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /comment/recipe/{id} [post]
func (ctl *CommentController) CommentByRecipe(c *gin.Context) {
	var comment models.CommentDTO
	idRecipe := c.Param("id")

	if utils.BindAndValidate(c, &comment) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(c, idRecipe, &comment)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// DeleteCommentByID
// @Summary Delete a comment by ID
// @Description Delete a comment from database
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID (UUID)" format(uuid)
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /comment/{id} [delete]
func (ctl *CommentController) DeleteComment(c *gin.Context) {
	idComment := c.Param("id")
	data, code, entity, errors := ctl.Repo.Delete(idComment)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
