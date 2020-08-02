package main

import (
	"log"

	"github.com/emmanuelneri/microservices-orchestration/commons/avro"
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

	codec, err := avro.LoadAvroCodec("apiRequestedSchema.avsc")
	if err != nil {
		panic(err)
	}

	processor.ApiRequestedSubscriber(processorConsumer, codec)

	select {}
}
