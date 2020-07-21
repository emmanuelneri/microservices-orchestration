package main

import (
	"github.com/emmanuelneri/microservices-orchestration/commonsconfig"
	"github.com/emmanuelneri/microservices-orchestration/processor/consumer"
	"log"
)

const (
	consumerGroupName = "processor-consumer-group"
	topic             = "ApiRequested"
)

func main() {
	bootstrapServers := commonsconfig.KafkaBootstrapServersFromEnvOrDefault()
	log.Print("Processor started")
	log.Print("Kafka bootstrapServers: ", bootstrapServers)

	go func() {
		consumer.SubscribeConsumer(bootstrapServers, consumerGroupName, topic)
	}()

	select {}
}
