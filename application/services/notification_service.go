package services

import (
	"errors"
	"student-crud/domain/events"
	"student-crud/domain/models"
	"student-crud/domain/ports"
)

type NotificationService struct {
	repo ports.NotificationRepository
}

func NewNotificationService(repo ports.NotificationRepository) *NotificationService {
	if repo == nil {
		panic("NotificationRepository cannot be nil")
	}
	return &NotificationService{repo: repo}
}

func (s *NotificationService) ProcessEvent(ev events.NotificationEvent) error {
	if ev.StudentID == 0 || ev.TeacherID == 0 || ev.Message == "" {
		return errors.New("invalid event")
	}

	notification := &models.Notification{
		StudentID: ev.StudentID,
		TeacherID: ev.TeacherID,
		Message:   ev.Message,
	}
	return s.repo.Save(notification)
}
