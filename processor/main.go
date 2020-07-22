package main

import (
	"log"

	"github.com/emmanuelneri/microservices-orchestration/processor/subscriber"
	"github.com/emmanuelneri/microservices-orchestration/processor/infra"
)

const (
	consumerGroupName = "processor-consumer-group"
)

func main() {
	log.Print("Processor started")
	processorConsumer := infra.CreateConsumer(consumerGroupName)

	go func() {
		subscriber.ApiRequestedSubscriber(processorConsumer)
	}()

	select {}
}
