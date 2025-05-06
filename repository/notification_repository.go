package repository

import (
	"fmt"
	"net/http"
	"service-pace11/config"
	"service-pace11/models"
	"service-pace11/utils"

	"github.com/gin-gonic/gin"
)

type NotificationRepository interface {
	Index(c *gin.Context) ([]models.Notification, int, any, int64, int, int)
	Save(c *gin.Context, userID string, likeByUserID string, typeNotif string, title string) bool
}

type notificationRepo struct{}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepo{}
}

func (r *notificationRepo) Index(c *gin.Context) ([]models.Notification, int, any, int64, int, int) {
	userIdRaw, exist := c.Get("user_id")
	var notifications []models.Notification
	var total int64

	if !exist {
		return nil, http.StatusUnauthorized, "access", 0, 0, 0
	}

	query := config.DB.Model(&models.Notification{})
	query = query.Where("user_id = ? ", userIdRaw)
	query = query.Order("created_at DESC")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&notifications)

	return notifications, http.StatusOK, "notification", total, page, limit
}

func (r *notificationRepo) Save(c *gin.Context, userID string, likeByUserID string, typeNotif string, title string) bool {
	if likeByUserID == userID {
		return false
	}

	username, exist := c.Get("name")
	var notification models.Notification

	notification.UserID = userID
	notification.Type = models.NotificationType(typeNotif)

	if !exist {
		return false
	}

	switch typeNotif {
	case string(models.CommentType):
		notification.Message = fmt.Sprintf("'%s' comment your recipe '%s'", username, title)
	case string(models.SaveType):
		notification.Message = fmt.Sprintf("'%s' save your recipe '%s'", username, title)
	default:
		notification.Message = fmt.Sprintf("'%s' like your recipe '%s'", username, title)
	}

	if err := config.DB.Save(&notification).Error; err != nil {
		return false
	}

	return true
}
