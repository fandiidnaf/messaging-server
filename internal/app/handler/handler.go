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

// @Summary      Register Token Device
// @Description  Register device token to FCM server
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        request body map[string]string true "Device Token"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/register-token [post]
func (h *NotificationHandler) RegisterTokenDevice(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RegisterTokenDevice(req.Token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "token registered successfully"})
}

// @Summary      Get Registered Tokens
// @Description  Get all registered device tokens
// @Tags         notification
// @Produce      json
// @Success      200  {array}  string
// @Router       /api/tokens [get]
func (h *NotificationHandler) GetRegisteredTokens(c *gin.Context) {
	tokens, err := h.service.GetRegisteredTokens()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
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
	case req.Tokens != nil:
		notifType = model.NotificationType{
			Tokens: req.Tokens,
		}
	}

	response, err := h.service.SendNotification(notifType, req.Title, req.Body, req.Data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Notification sent with id: %v", response)})
}

// @Summary      Subscribe To Topic
// @Description  Subscribe devices to a topic
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        request body map[string]interface{} true "Tokens and Topic"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/subscribe [post]
func (h *NotificationHandler) SubscribeToTopic(c *gin.Context) {
	var req struct {
		Tokens []string `json:"tokens" binding:"required"`
		Topic  string   `json:"topic" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.SubscribeToTopic(req.Tokens, req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res})
}

// @Summary      Unsubscribe From Topic
// @Description  Unsubscribe devices from a topic
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        request body map[string]interface{} true "Tokens and Topic"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/unsubscribe [post]
func (h *NotificationHandler) UnsubscribeFromTopic(c *gin.Context) {
	var req struct {
		Tokens []string `json:"tokens" binding:"required"`
		Topic  string   `json:"topic" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.UnsubscribeFromTopic(req.Tokens, req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res})
}
