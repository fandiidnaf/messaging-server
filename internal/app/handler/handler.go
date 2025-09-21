package handler

import (
	"fmt"
	"net/http"

	"github.com/fandiidnaf/messaging-server/internal/app/model"
	"github.com/fandiidnaf/messaging-server/internal/app/service"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service *service.NotificationService
}

func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

// @Summary      Send Notification
// @Description  Send Notification to device with FCM
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        request body model.NotificationRequest true "Notification Request"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Router       /api/notify [post]
func (h *NotificationHandler) SendNotification(c *gin.Context) {
	req := model.NotificationRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notifType model.NotificationType

	switch {
	case req.Token != "":
		notifType = model.NotificationType{
			Token: req.Token,
		}
	case req.Topic != "":
		notifType = model.NotificationType{
			Topic: req.Topic,
		}
	case req.Condition != "":
		notifType = model.NotificationType{
			Condition: req.Condition,
		}
	}

	response, err := h.service.SendNotification(notifType, req.Title, req.Body, req.Data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Notification sent with id: %v", response)})
}
