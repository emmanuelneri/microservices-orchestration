package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emmanuelneri/microservices-orchestration/api/structs"
	kafkaProducer "github.com/emmanuelneri/microservices-orchestration/commons/kafka"
	_ "github.com/linkedin/goavro/v2"
)

type RequestHandler struct {
	Producer *kafkaProducer.Producer
}

func (requestHandler *RequestHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBody structs.RequestBody
	bodyError := json.NewDecoder(request.Body).Decode(&requestBody)

	if bodyError != nil {
		http.Error(responseWriter, "invalid request. error: "+bodyError.Error(), http.StatusBadRequest)
		return
	}

	log.Println("API requested: ", requestBody)

	producer := requestHandler.Producer
	key := []byte(requestBody.Identifier)
	produceError := producer.Produce(key, requestBody.ToMap())

	if produceError != nil {
		http.Error(responseWriter, "internal error", http.StatusBadRequest)
		log.Panicln("producer error", produceError)
		return
	}

	responseWriter.WriteHeader(http.StatusAccepted)
}
