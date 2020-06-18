package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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
