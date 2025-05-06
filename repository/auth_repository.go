package repository

import (
	"net/http"
	"service-pace11/config"
	"service-pace11/models"
	"service-pace11/utils"
	"time"
)

type AuthRepository interface {
	Login(login *models.LoginDTO) (any, int, string, map[string]string)
	Register(register *models.RegisterDTO) (any, int, string, map[string]string)
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) Login(login *models.LoginDTO) (any, int, string, map[string]string) {
	var user models.User

	if err := config.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		return nil, http.StatusNotFound, "user", nil
	}

	if !utils.CheckPasswordHash(login.Password, user.Password) {
		return nil, http.StatusUnauthorized, "invalid credentials", nil
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, http.StatusInternalServerError, "generate token", nil
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response := map[string]any{
		"token":   token,
		"expires": time.Now().Add(time.Hour * 24).UTC().Format(time.RFC3339),
		"user":    userResponse,
	}

	return response, http.StatusOK, "user", nil
}

func (r *authRepository) Register(register *models.RegisterDTO) (any, int, string, map[string]string) {
	hashed, err := utils.HashPassword(register.Password)

	if err != nil {
		return nil, http.StatusInternalServerError, "hash password", nil
	}

	registerPayload := models.User{
		Name:     register.Name,
		Address:  &register.Address,
		Email:    register.Email,
		Password: hashed,
	}

	if err := config.DB.Save(&registerPayload).Error; err != nil {
		return nil, http.StatusInternalServerError, "user", map[string]string{
			"message": err.Error(),
		}
	}

	return registerPayload, http.StatusCreated, "user", nil
}
