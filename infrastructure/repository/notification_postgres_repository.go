package repository

import (
	"student-crud/config"
	"student-crud/domain/models"
)

type PostgresNotificationRepository struct{}

func NewPostgresNotificationRepository() *PostgresNotificationRepository {
	return &PostgresNotificationRepository{}
}

func (r *PostgresNotificationRepository) Save(n *models.Notification) error {
	return config.DB.Create(n).Error
}
