package main

import (
	"log"
	"net/http"

	"github.com/emmanuelneri/microservices-orchestration/sync-receiver/internal/handler"
)

func main() {
	log.Print("Sync receiver started")

	requestHandler := handler.NewRequestHandler()
	http.HandleFunc("/", requestHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
