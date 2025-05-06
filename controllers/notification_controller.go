package controllers

import (
	"service-pace11/repository"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	Repo repository.NotificationRepository
}

func NewNotificationController(repo repository.NotificationRepository) *NotificationController {
	return &NotificationController{Repo: repo}
}

// GetNotifications
// @Summary Get list of notifications
// @Description Retrieve all notifications with pagination
// @Tags Notifications
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} utils.PaginatedResponses
// @Security BearerAuth
// @Router /notifications [get]
func (ctl *NotificationController) GetNotifications(c *gin.Context) {
	data, code, entity, total, page, limit := ctl.Repo.Index(c)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}
