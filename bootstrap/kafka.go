package bootstrap

import (
	"student-crud/application/services"
	"student-crud/infrastructure/kafka"
	"student-crud/infrastructure/repository"
)

func InitKafkaConsumers() {
	notifRepo := repository.NewPostgresNotificationRepository()
	notifService := services.NewNotificationService(notifRepo)
	notifConsumer := kafka.NewNotificationConsumer(notifService)
	go notifConsumer.Start("student_notifications", []string{"localhost:9092"})
}
