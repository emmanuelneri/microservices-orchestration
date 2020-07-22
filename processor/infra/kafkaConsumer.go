package infra

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelneri/microservices-orchestration/commonsconfig"
)

func CreateConsumer(consumerGroupName string) *kafka.Consumer {
	bootstrapServers := commonsconfig.KafkaBootstrapServersFromEnvOrDefault()

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
