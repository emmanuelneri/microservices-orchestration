package main

import (
	"log"
	"net/http"

	"github.com/emmanuelneri/microservices-orchestration/api/handler"
	"github.com/emmanuelneri/microservices-orchestration/api/infra"
	"github.com/emmanuelneri/microservices-orchestration/commons/avro"
	_ "github.com/linkedin/goavro/v2"
)

const (
	topic = "ApiRequested"
)

func main() {
	log.Print("API started")
	producer := infra.CreateKafkaProducer()

	codec, err := avro.LoadAvroCodec("apiRequestedSchema.avsc")
	if err != nil {
		panic(err)
	}

	requestHandler := handler.CreateRequestHandler(producer, topic, codec)
	http.HandleFunc("/", requestHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
