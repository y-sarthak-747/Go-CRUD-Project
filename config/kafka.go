package config

import (
	"log"
	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer

func InitKafkaProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	var err error
	Producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	log.Println("✅ Kafka Producer initialized")
}
