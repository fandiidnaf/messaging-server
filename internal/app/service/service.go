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

func (s *NotificationService) SendNotification(notificationType model.NotificationType, title, body string, data map[string]string) (string, error) {
	return s.fcm.SendNotification(notificationType, title, body, data)
}
