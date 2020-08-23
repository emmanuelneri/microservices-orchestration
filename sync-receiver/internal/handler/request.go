package handler

import (
	"encoding/json"
	"github.com/emmanuelneri/microservices-orchestration/sync-receiver/internal/service"
	"github.com/emmanuelneri/microservices-orchestration/sync-receiver/pkg/processor"
	"github.com/emmanuelneri/microservices-orchestration/sync-receiver/pkg/receiver"
	"log"
	"net/http"
)

type RequestHandler interface {
	Handle(responseWriter http.ResponseWriter, request *http.Request)
}

type RequestHandlerImpl struct {
	service service.ProcessorService
}

func NewRequestHandler() RequestHandler {
	return &RequestHandlerImpl{service: service.NewRequestHandler()}
}

func (requestHandler *RequestHandlerImpl) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBody receiver.Body
	bodyError := json.NewDecoder(request.Body).Decode(&requestBody)

	if bodyError != nil {
		http.Error(responseWriter, "invalid request. error: "+bodyError.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Sync receiver requested: ", requestBody)

	processorBody := processor.Body{
		Identifier: requestBody.Identifier,
		Customer:   requestBody.Customer,
	}

	response, err := requestHandler.service.Post(processorBody)
	if err != nil {
		if response == nil {
			http.Error(responseWriter, "Internal error", http.StatusBadRequest)
			return
		}

		http.Error(responseWriter, err.Error(), response.StatusCode)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
