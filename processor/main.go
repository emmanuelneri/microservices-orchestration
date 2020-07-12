package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func main() {
	bootstrapServers := getBootstrapServers()
	log.Print("Processor started")
	log.Print("Kafka bootstrapServers: ", bootstrapServers)

	consumer, consumerError := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          "processor-consumer-group",
		"auto.offset.reset": "earliest",
	})

	if consumerError != nil {
		log.Fatal("Failed to create consumer.", consumerError)
	}

	topic := "ApiRequested"
	subscriberError := consumer.Subscribe(topic, nil)

	if subscriberError != nil {
		log.Fatal("Failed to subscribe topic: ", topic, subscriberError)
	}
	var run = true
	for run == true {
		eventConsumed := consumer.Poll(0)
		switch event := eventConsumed.(type) {
		case *kafka.Message:
			fmt.Printf("%% Message Consumed. partition: [%d] - key: %s - value: %s\n", event.TopicPartition.Partition, string(event.Key), string(event.Value))
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", event)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", event)
			run = false
		default:

		}
	}

	consumer.Close()
}

// TODO Configmap
func getBootstrapServers() string {
	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if bootstrapServers == "" {
		return "localhost:9092"
	}

	return bootstrapServers
}
