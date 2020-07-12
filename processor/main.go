package main

import (
	"log"
	"os"
	"processormodule/consumer"
)

const (
	consumerGroupName = "processor-consumer-group"
	topic 			  = "ApiRequested"
)

func main() {
	bootstrapServers := getBootstrapServers()
	log.Print("Processor started")
	log.Print("Kafka bootstrapServers: ", bootstrapServers)


	go func() {
		consumer.SubscribeConsumer(bootstrapServers, consumerGroupName, topic)
	}()

	select {}
}

func getBootstrapServers() string {
	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if bootstrapServers == "" {
		return "localhost:9092"
	}

	return bootstrapServers
}
