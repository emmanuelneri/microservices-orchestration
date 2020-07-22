package main

import (
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelneri/microservices-orchestration/api/handler"
	"github.com/emmanuelneri/microservices-orchestration/commonsconfig"
)

const (
	topic = "ApiRequested"
)

func main() {
	log.Print("API started")
	producer := createKafkaProducer()

	requestHandler := &handler.RequestHandler{KafkaProducer: producer, Topic: topic}
	http.HandleFunc("/", requestHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createKafkaProducer() *kafka.Producer {
	bootstrapServers := commonsconfig.KafkaBootstrapServersFromEnvOrDefault()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	log.Print("Kafka Producer started. bootstrapServers: ", bootstrapServers)

	return producer
}
