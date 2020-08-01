package main

import (
	"github.com/linkedin/goavro/v2"
	"io/ioutil"
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

	schema, err := ioutil.ReadFile("apiRequestedSchema.avsc")
	if err != nil {
		panic(err)
	}

	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		panic(err)
	}

	go func() {
		subscriber.ApiRequestedSubscriber(processorConsumer, codec)
	}()

	select {}
}
