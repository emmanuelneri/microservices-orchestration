package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelneri/microservices-orchestration/api/structs"
	kafkaProducer "github.com/emmanuelneri/microservices-orchestration/commons/kafka"
	avro "github.com/linkedin/goavro/v2"
)

// RequestHandler use to get db address for HTTP Handle method set
type RequestHandler struct {
	kafkaProducer *kafka.Producer
	topic         string
	codec         *avro.Codec
	deliveryChan  chan kafka.Event
}

func CreateRequestHandler(producer *kafka.Producer, topic string, codec *avro.Codec) *RequestHandler {
	return &RequestHandler{kafkaProducer: producer, topic: topic, codec: codec, deliveryChan: make(chan kafka.Event, 10000)}
}

func (requestHandler *RequestHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBody structs.RequestBody
	bodyError := json.NewDecoder(request.Body).Decode(&requestBody)

	if bodyError != nil {
		http.Error(responseWriter, "invalid request. error: "+bodyError.Error(), http.StatusBadRequest)
		return
	}

	log.Println("API requested: ", requestBody)

	produceError := produce(requestBody, requestHandler.kafkaProducer, requestHandler.topic, requestHandler.codec, requestHandler.deliveryChan)

	if produceError != nil {
		http.Error(responseWriter, "internal error", http.StatusBadRequest)
		log.Panicln("producer error", produceError)
		return
	}

	responseWriter.WriteHeader(http.StatusAccepted)
}

func produce(requestBody structs.RequestBody, producer *kafka.Producer, topic string, codec *avro.Codec, deliveryChan chan kafka.Event) error {
	key := []byte(requestBody.Identifier)

	binary, err := codec.BinaryFromNative(nil, requestBody.ToMap())
	if err != nil {
		panic(err)
	}

	return kafkaProducer.ProduceMessage(key, binary, producer, topic, deliveryChan)
}
