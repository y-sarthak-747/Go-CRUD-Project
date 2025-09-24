package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"student-crud/config"
)

func SendMessage(topic string, msg []byte) error {
	if config.Producer == nil {
		return fmt.Errorf("Kafka producer not initialized")
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}
	_, _, err := config.Producer.SendMessage(message)
	return err
}
