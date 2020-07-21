package main

import (
	"log"
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelneri/microservices-orchestration/api/handler"
	"github.com/emmanuelneri/microservices-orchestration/commonsconfig"
)

const (
	topic = "ApiRequested"
)

func main() {
	bootstrapServers := commonsconfig.KafkaBootstrapServersFromEnvOrDefault()
	log.Print("API started")
	log.Print("Kafka bootstrapServers: ", bootstrapServers)

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	requestHandler := &handler.RequestHandler{KafkaProducer: producer, Topic: topic}
	http.HandleFunc("/", requestHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBootstrapServers() string {
	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if bootstrapServers == "" {
		return "localhost:9092"
	}

	return bootstrapServers
}
