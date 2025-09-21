package service

import (
	"github.com/fandiidnaf/messaging-server/internal/app/model"
	"github.com/fandiidnaf/messaging-server/internal/app/repository"
)

type NotificationService struct {
	fcm repository.FCMRepository
}

func NewNotificationService(fcm repository.FCMRepository) *NotificationService {
	return &NotificationService{fcm: fcm}
}

func (s *NotificationService) RegisterTokenDevice(token string) error {
	return s.fcm.RegisterTokenDevice(token)
}

func (s *NotificationService) GetRegisteredTokens() ([]string, error) {
	return s.fcm.GetRegisteredTokens()
}

func (s *NotificationService) SendNotification(notificationType model.NotificationType, title, body string, data map[string]string) (string, error) {
	return s.fcm.SendNotification(notificationType, title, body, data)
}

func (s *NotificationService) SubscribeToTopic(tokens []string, topic string) (string, error) {
	return s.fcm.SubscribeToTopic(tokens, topic)
}

func (s *NotificationService) UnsubscribeFromTopic(tokens []string, topic string) (string, error) {
	return s.fcm.UnsubscribeFromTopic(tokens, topic)
}
