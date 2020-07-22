package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelneri/microservices-orchestration/api/infra"
	"github.com/emmanuelneri/microservices-orchestration/api/structs"
)

// RequestHandler use to get db address for HTTP Handle method set
type RequestHandler struct {
	kafkaProducer *kafka.Producer
	topic         string
	deliveryChan  chan kafka.Event
}

func CreateRequestHandler(producer *kafka.Producer, topic string) *RequestHandler {
	return &RequestHandler{kafkaProducer: producer, topic: topic, deliveryChan: make(chan kafka.Event, 10000)}
}

func (requestHandler *RequestHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBodyAsJson structs.RequestBody
	bodyError := json.NewDecoder(request.Body).Decode(&requestBodyAsJson)

	if bodyError != nil {
		http.Error(responseWriter, "invalid request. error: "+bodyError.Error(), http.StatusBadRequest)
		return
	}

	log.Println("API requested: ", requestBodyAsJson)

	produceError := produce(requestBodyAsJson, requestHandler.kafkaProducer, requestHandler.topic, requestHandler.deliveryChan)

	if produceError != nil {
		http.Error(responseWriter, "internal error", http.StatusBadRequest)
		log.Panicln("producer error", produceError)
		return
	}

	responseWriter.WriteHeader(http.StatusAccepted)
}

func produce(requestBodyAsJson structs.RequestBody, producer *kafka.Producer, topic string, deliveryChan chan kafka.Event) error {
	key := []byte(requestBodyAsJson.Identifier)
	value := new(bytes.Buffer)
	json.NewEncoder(value).Encode(&requestBodyAsJson)

	return infra.ProduceMessage(key, value.Bytes(), producer, topic, deliveryChan)
}
