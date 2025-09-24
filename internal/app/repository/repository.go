package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"maps"

	"firebase.google.com/go/v4/messaging"
	"github.com/fandiidnaf/messaging-server/config"
	"github.com/fandiidnaf/messaging-server/internal/app/model"
)

type FCMRepository interface {
	RegisterTokenDevice(token string) error
	GetRegisteredTokens() ([]string, error)
	SendNotification(notificationType model.NotificationType, title, body string, data map[string]string) (string, error)
	SubscribeToTopic(tokens []string, topic string) (string, error)
	UnsubscribeFromTopic(tokens []string, topic string) (string, error)
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

var tokens []string = make([]string, 0)

func (r *fcmRepository) RegisterTokenDevice(token string) error {
	if token == "" {
		return errors.New("The token cannot empty")
	}

	tokens = append(tokens, token)

	return nil
}

func (r *fcmRepository) GetRegisteredTokens() ([]string, error) {
	return tokens, nil
}

func (r *fcmRepository) SendNotification(notificationType model.NotificationType, title, body string, data map[string]string) (string, error) {

	mergedData := map[string]string{
		"title": title,
		"body":  body,
	}

	maps.Copy(mergedData, data)

	if notificationType.Tokens != nil {
		msgs := &messaging.MulticastMessage{
			Tokens: notificationType.Tokens,
			Data:   mergedData,
			Android: &messaging.AndroidConfig{
				Priority: "high",
			},
		}

		response, err := r.client.SendEachForMulticast(context.Background(), msgs)

		if err != nil {
			return "", err
		}

		if response.FailureCount > 0 {
			failedTokens := make([]string, 0)

			for idx, resp := range response.Responses {
				if !resp.Success {
					failedTokens = append(failedTokens, tokens[idx])
				}
			}

			return "", fmt.Errorf("Failed to send notification to %d tokens", len(failedTokens))
		}

		return "messages sent successfully", nil
	}

	// var duration = time.Second * 0

	msg := &messaging.Message{
		Token:     notificationType.Token,
		Topic:     notificationType.Topic,
		Condition: notificationType.Condition,
		Android: &messaging.AndroidConfig{
			Priority: "high",
			// TTL:      &duration,
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "5",
			},
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

func (r *fcmRepository) SubscribeToTopic(tokens []string, topic string) (string, error) {
	if tokens == nil {
		return "", errors.New("tokens cannot be empty")
	}

	if topic == "" {
		return "", errors.New("topic cannot be empty")
	}

	_, err := r.client.SubscribeToTopic(context.Background(), tokens, topic)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("subscribed successfully to topic: %s", topic), nil
}

func (r *fcmRepository) UnsubscribeFromTopic(tokens []string, topic string) (string, error) {
	if tokens == nil && len(tokens) == 0 {
		return "", errors.New("tokens cannot be nil or empty")
	}

	if topic == "" {
		return "", errors.New("topic cannot be empty")
	}

	_, err := r.client.UnsubscribeFromTopic(context.Background(), tokens, topic)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("unsubscribed successfully from topic: %s", topic), nil
}
