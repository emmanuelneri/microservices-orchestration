package main

import (
	"github.com/emmanuelneri/microservices-orchestration/sync-processor/internal/handler"
	"log"
	"net/http"
)

func main() {
	log.Print("Sync processor started")

	processortHandler := handler.NewRequestHandler()
	http.HandleFunc("/", processortHandler.Handle)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
