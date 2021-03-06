package main

import (
	"log"
	"net/http"

	"github.com/emmanuelneri/microservices-orchestration/api/handler"
	kafkaProducer "github.com/emmanuelneri/microservices-orchestration/commons/kafka"
	_ "github.com/linkedin/goavro/v2"
)

const (
	topic = "ApiRequested"
)

func main() {
	log.Print("API started")

	producer, err := kafkaProducer.CreateProducer(topic, "apiRequestedSchema.avsc")
	if err != nil {
		panic(err)
	}

	requestHandler := handler.RequestHandler{Producer: producer}
	http.HandleFunc("/", requestHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
