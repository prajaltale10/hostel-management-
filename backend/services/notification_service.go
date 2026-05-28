package services

import (
	"context"
	"log"

	"hostel-saas/models"
	"gorm.io/gorm"
)

type NotificationProvider interface {
	Send(recipient string, message string) error
}

type MockEmailProvider struct{}

func (p *MockEmailProvider) Send(recipient string, message string) error {
	log.Printf("[Mock Email] Sending to %s: %s", recipient, message)
	return nil
}

type MockSMSProvider struct{}

func (p *MockSMSProvider) Send(recipient string, message string) error {
	log.Printf("[Mock SMS] Sending to %s: %s", recipient, message)
	return nil
}

type NotificationService interface {
	SendNotification(ctx context.Context, orgID uint, studentID uint, channel models.NotificationChannel, recipient, message string) error
}

type notificationService struct {
	db       *gorm.DB
	emailPvd NotificationProvider
	smsPvd   NotificationProvider
}

func NewNotificationService(db *gorm.DB) NotificationService {
	return &notificationService{
		db:       db,
		emailPvd: &MockEmailProvider{},
		smsPvd:   &MockSMSProvider{},
	}
}

func (s *notificationService) SendNotification(ctx context.Context, orgID uint, studentID uint, channel models.NotificationChannel, recipient, message string) error {
	var err error
	if channel == models.ChannelEmail {
		err = s.emailPvd.Send(recipient, message)
	} else {
		err = s.smsPvd.Send(recipient, message)
	}

	status := models.NotifStatusSent
	errorMsg := ""
	if err != nil {
		status = models.NotifStatusFailed
		errorMsg = err.Error()
	}

	logEntry := &models.NotificationLog{
		OrgID:     orgID,
		StudentID: studentID,
		Channel:   channel,
		Recipient: recipient,
		Message:   message,
		Status:    status,
		ErrorMsg:  errorMsg,
	}

	return s.db.WithContext(ctx).Create(logEntry).Error
}
