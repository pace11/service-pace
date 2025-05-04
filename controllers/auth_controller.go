package controllers

import (
	"service-pace11/models"
	"service-pace11/repository"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Repo repository.AuthRepository
}

func NewAuthController(repo repository.AuthRepository) *AuthController {
	return &AuthController{Repo: repo}
}

// AuthLogin
// @Summary Login User
// @Description Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body models.LoginDTO true "Login payload"
// @Success 201 {object} utils.StandardResponses
// @Router /auth/login [post]
func (ctl *AuthController) Login(c *gin.Context) {
	var payload models.LoginDTO

	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Login(&payload)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

// AuthRegister
// @Summary Register User
// @Description Register User
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body models.RegisterDTO true "Register payload"
// @Success 201 {object} utils.StandardResponses
// @Router /auth/register [post]
func (ctl *AuthController) Register(c *gin.Context) {
	var payload models.RegisterDTO

	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Register(&payload)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
