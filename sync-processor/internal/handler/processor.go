package handler

import (
	"encoding/json"
	"github.com/emmanuelneri/microservices-orchestration/sync-processor/pkg/processor"
	"log"
	"net/http"
)

type RequestHandler interface {
	Handle(responseWriter http.ResponseWriter, request *http.Request)
}

type RequestHandlerImpl struct {
}

func NewRequestHandler() RequestHandler {
	return &RequestHandlerImpl{}
}

func (requestHandler *RequestHandlerImpl) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	var body processor.Body
	bodyError := json.NewDecoder(request.Body).Decode(&body)

	if bodyError != nil {
		http.Error(responseWriter, "invalid request. error: "+bodyError.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Sync processor requested: ", body)
	responseWriter.WriteHeader(http.StatusOK)
}
