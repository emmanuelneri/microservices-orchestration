package main

import (
	"log"

	"github.com/emmanuelneri/microservices-orchestration/commons/avro"
	kafkaConsumer "github.com/emmanuelneri/microservices-orchestration/commons/kafka"
	"github.com/emmanuelneri/microservices-orchestration/processor/subscriber"
	_ "github.com/linkedin/goavro/v2"
)

const (
	consumerGroupName = "processor-consumer-group"
)

func main() {
	log.Print("Processor started")
	processorConsumer := kafkaConsumer.CreateConsumer(consumerGroupName)

	codec, err := avro.LoadAvroCodec("apiRequestedSchema.avsc")
	if err != nil {
		panic(err)
	}

	go func() {
		subscriber.ApiRequestedSubscriber(processorConsumer, codec)
	}()

	select {}
}
