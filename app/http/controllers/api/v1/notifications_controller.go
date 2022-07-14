package v1

import (
	"gohub/app/models/notification"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type NotificationsController struct {
	BaseAPIController
}

func (ctrl *NotificationsController) Index(c *gin.Context) {
	notifications := notification.All()
	response.SuccessWithData(c, gin.H{
		"data": notifications,
	})
}

func (ctrl *NotificationsController) Show(c *gin.Context) {
	notificationModel := notification.Get(c.Param("id"))
	if notificationModel.ID == 0 {
		response.NotFound(c)
		return
	}
	response.SuccessWithData(c, gin.H{
		"data": notificationModel,
	})
}

func (ctrl *NotificationsController) GetNextQueued(c *gin.Context) {
	notificationModel := notification.GetOldestUnread()
	if notificationModel.ID == 0 {
		response.NotFound(c)
		return
	}
	response.SuccessWithData(c, gin.H{
		"data": notificationModel,
	})
	notificationModel.MarkAsRead()
}
