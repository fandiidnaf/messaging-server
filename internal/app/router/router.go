package router

import (
	"github.com/fandiidnaf/messaging-server/internal/app/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(notificationHandler *handler.NotificationHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// Register & get tokens
		api.POST("/register-token", notificationHandler.RegisterTokenDevice)
		api.GET("/tokens", notificationHandler.GetRegisteredTokens)

		// Notifications
		api.POST("/notify", notificationHandler.SendNotification)

		// Topic management
		api.POST("/subscribe", notificationHandler.SubscribeToTopic)
		api.POST("/unsubscribe", notificationHandler.UnsubscribeFromTopic)
	}

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
