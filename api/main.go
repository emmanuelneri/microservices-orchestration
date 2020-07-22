package main

import (
	"log"
	"net/http"

	"github.com/emmanuelneri/microservices-orchestration/api/handler"
	"github.com/emmanuelneri/microservices-orchestration/api/infra"
)

const (
	topic = "ApiRequested"
)

func main() {
	log.Print("API started")
	producer := infra.CreateKafkaProducer()

	requestHandler := handler.CreateRequestHandler(producer, topic)
	http.HandleFunc("/", requestHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
