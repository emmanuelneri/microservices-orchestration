package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"net/http"
	"os"

	"apimodule/handler"
)

const (
	topic = "ApiRequested"
)

func main() {
	bootstrapServers := getBootstrapServers()
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
