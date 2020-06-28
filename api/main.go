package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Print("API started")

	topic := "myTopic"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		log.Print("API requested")

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic},
			Value:          body,
		}, nil)

		if err != nil {
			log.Fatalf("Failed to produce: %s\n", err)

		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
