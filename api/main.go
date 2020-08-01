package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/linkedin/goavro/v2"
	"github.com/emmanuelneri/microservices-orchestration/api/handler"
	"github.com/emmanuelneri/microservices-orchestration/api/infra"
)

const (
	topic = "ApiRequested"
)

func main() {
	log.Print("API started")
	producer := infra.CreateKafkaProducer()

	schema, err := ioutil.ReadFile("apiRequestedSchema.avsc")
	if err != nil {
		panic(err)
	}

	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		panic(err)
	}

	requestHandler := handler.CreateRequestHandler(producer, topic, codec)
	http.HandleFunc("/", requestHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
