package main

import (
	"log"

	_ "github.com/fandiidnaf/messaging-server/docs"
	"github.com/fandiidnaf/messaging-server/internal/app/handler"
	"github.com/fandiidnaf/messaging-server/internal/app/repository"
	"github.com/fandiidnaf/messaging-server/internal/app/router"
	"github.com/fandiidnaf/messaging-server/internal/app/service"
)

func main() {
	// layer repository
	fcmRepo := repository.NewFCMRepository()

	// layer service
	notificationService := service.NewNotificationService(fcmRepo)

	// layer handler
	notificationHandler := handler.NewNotificationHandler(notificationService)

	r := router.SetupRouter(notificationHandler)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}

}
