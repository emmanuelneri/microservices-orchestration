package main

import (
	"log"

	kafkaConsumer "github.com/emmanuelneri/microservices-orchestration/commons/kafka"
	"github.com/emmanuelneri/microservices-orchestration/processor/processor"
	_ "github.com/linkedin/goavro/v2"
)

const (
	consumerGroupName = "processor-consumer-group"
)

func main() {
	log.Print("Processor started")
	processorConsumer := kafkaConsumer.CreateConsumer(consumerGroupName)
	processor.ApiRequestedSubscriber(processorConsumer)

	select {}
}
