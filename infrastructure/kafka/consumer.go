package kafka

import (
	"encoding/json"
	"log"
	"student-crud/application/services"
	"student-crud/domain/events"

	"github.com/IBM/sarama"
)

type NotificationConsumer struct {
	service *services.NotificationService
}

func NewNotificationConsumer(svc *services.NotificationService) *NotificationConsumer {
	return &NotificationConsumer{service: svc}
}

func (c *NotificationConsumer) Start(topic string, brokers []string) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	pc, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error starting consumer: %v", err)
	}
	log.Printf("✅ Kafka consumer started on topic %s", topic)

	go func() {
		for msg := range pc.Messages() {
			var ev events.NotificationEvent
			if err := json.Unmarshal(msg.Value, &ev); err != nil {
				log.Printf("Invalid event JSON: %v", err)
				continue
			}
			if err := c.service.ProcessEvent(ev); err != nil {
				log.Printf("Failed to process event: %v", err)
				continue
			}
			log.Printf("✅ Processed notification: %+v", ev)
		}
	}()
}
