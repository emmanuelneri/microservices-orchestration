package kafka

import (
	"log"

	"github.com/emmanuelneri/microservices-orchestration/commons/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func CreateConsumer(consumerGroupName string) *kafka.Consumer {
	bootstrapServers := config.KafkaBootstrapServersFromEnvOrDefault()

	consumer, consumerError := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          consumerGroupName,
		"auto.offset.reset": "earliest",
	})

	if consumerError != nil {
		log.Fatal("Failed to create consumer.", consumerError)
	}

	log.Print("Kafka Consumer started. bootstrapServers: ", bootstrapServers)

	return consumer
}
