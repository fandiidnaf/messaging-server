package repository

import (
	"context"
	"log"
	"maps"

	"firebase.google.com/go/v4/messaging"
	"github.com/fandiidnaf/messaging-server/config"
	"github.com/fandiidnaf/messaging-server/internal/app/model"
)

type FCMRepository interface {
	SendNotification(notificationType model.NotificationType, title, body string, data map[string]string) (string, error)
}

type fcmRepository struct {
	client *messaging.Client
}

func NewFCMRepository() FCMRepository {
	app := config.InitFirebaseApp()

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	return &fcmRepository{client: client}
}

func (r *fcmRepository) SendNotification(notificationType model.NotificationType, title, body string, data map[string]string) (string, error) {

	mergedData := map[string]string{
		"title": title,
		"body":  body,
	}

	maps.Copy(mergedData, data)

	msg := &messaging.Message{
		Token:     notificationType.Token,
		Topic:     notificationType.Topic,
		Condition: notificationType.Condition,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
		// Notification: &messaging.Notification{
		// 	Title:    title,
		// 	Body:     body,
		// 	ImageURL: "https://tse1.mm.bing.net/th/id/OIP.OJKlDIGrja5MQ-yvqHhASAAAAA?rs=1&pid=ImgDetMain&o=7&rm=3",
		// },
		Data: mergedData,
	}

	return r.client.Send(context.Background(), msg)
}
