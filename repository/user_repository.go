package repository

import (
	"net/http"
	"service-pace11/config"
	"service-pace11/models"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetMe(c *gin.Context) (*models.UserDetailResponse, int, string, map[string]string)
	Update(c *gin.Context, user *models.UserUpdateDTO) (any, int, string, map[string]string)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetMe(c *gin.Context) (*models.UserDetailResponse, int, string, map[string]string) {
	userIDVal, exist := c.Get("user_id")
	var userDetail models.User

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	if err := config.DB.First(&userDetail, userIDVal).Error; err != nil {
		return nil, http.StatusNotFound, "user", nil
	}

	response := &models.UserDetailResponse{
		ID:      userDetail.ID,
		Name:    userDetail.Name,
		Email:   userDetail.Email,
		Address: *userDetail.Address,
	}

	return response, http.StatusOK, "user", nil
}

func (r *userRepository) Update(c *gin.Context, user *models.UserUpdateDTO) (any, int, string, map[string]string) {
	userIDVal, exist := c.Get("user_id")
	var existing models.User

	if !exist {
		return nil, http.StatusUnauthorized, "access", nil
	}

	if err := config.DB.First(&existing, userIDVal).Error; err != nil {
		return nil, http.StatusNotFound, "user", nil
	}

	if user.Password != "" {
		hashed, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, http.StatusInternalServerError, "user", nil
		}
		existing.Password = hashed
	}

	if user.Address != "" {
		existing.Address = &user.Address
	}

	existing.Name = user.Name
	existing.Email = user.Email

	if err := config.DB.Save(&existing).Error; err != nil {
		return nil, http.StatusInternalServerError, "user", map[string]string{
			"message": err.Error(),
		}
	}

	return existing, http.StatusOK, "user", nil
}
