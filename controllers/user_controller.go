package controllers

import (
	"service-pace11/models"
	"service-pace11/repository"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

// GetUserDetail
// @Summary Get a user data
// @Description Get detail user from jwt token
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /user/me [get]
func (ctl *UserController) GetMe(c *gin.Context) {
	data, code, entity, errors := ctl.Repo.GetMe(c)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// UpdateUser
// @Summary Update user data
// @Description Update user data from jwt token
// @Tags User
// @Accept json
// @Produce json
// @Param payload body models.UserUpdateDTO true "User payload"
// @Success 200 {object} utils.StandardResponses
// @Security BearerAuth
// @Router /user/update [get]
func (ctl *UserController) Update(c *gin.Context) {
	var user models.UserUpdateDTO

	if utils.BindAndValidate(c, &user) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(c, &user)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
