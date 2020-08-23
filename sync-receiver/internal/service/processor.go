package service

import (
	"bytes"
	"encoding/json"
	"github.com/emmanuelneri/microservices-orchestration/commons/config"
	"github.com/emmanuelneri/microservices-orchestration/sync-receiver/pkg/processor"
	"net/http"
)

const contentType = "application/json"

type ProcessorService interface {
	Post(processorBody processor.Body) (*http.Response, error)
}

type ProcessorServiceImpl struct {
	httpClient http.Client
	url        string
}

func NewRequestHandler() ProcessorService {
	return &ProcessorServiceImpl{httpClient: http.Client{}, url: config.SyncProcessorAppUrlFromEnvOrDefault()}
}

func (p *ProcessorServiceImpl) Post(processorBody processor.Body) (*http.Response, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(processorBody)
	if err != nil {
		return nil, err
	}

	return p.httpClient.Post(p.url, contentType, body)
}
