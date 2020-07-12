package handler

import (
	"apimodule/structs"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"net/http"
)

// RequestHandler use to get db address for HTTP Handle method set
type RequestHandler struct {
	KafkaProducer *kafka.Producer
	Topic 		  string
}

func (requestHandler *RequestHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBodyAsJson structs.RequestBody
	bodyError := json.NewDecoder(request.Body).Decode(&requestBodyAsJson)

	if bodyError != nil {
		http.Error(responseWriter, "invalid request. error: " + bodyError.Error(), http.StatusBadRequest)
		return
	}

	log.Println("API requested: ", requestBodyAsJson)

	deliveryChan := make(chan kafka.Event, 10000)
	produceError := produceMessage(requestBodyAsJson, requestHandler.KafkaProducer, requestHandler.Topic, deliveryChan)

	if(produceError != nil) {
		http.Error(responseWriter, "internal error", http.StatusBadRequest)
		log.Panicln("producer error", produceError)
		return
	}

	responseWriter.WriteHeader(http.StatusAccepted)
}

func produceMessage(requestBodyAsJson structs.RequestBody, producer *kafka.Producer, topic string, deliveryChan chan kafka.Event) error {
	requestBodyAsBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyAsBytes).Encode(&requestBodyAsJson)

	produceError := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(requestBodyAsJson.Identifier),
		Value:          requestBodyAsBytes.Bytes(),
	}, deliveryChan)

	if produceError != nil {
		log.Panicln("error to produce: ", produceError)
	}

	logProduceMessage(deliveryChan)

	return produceError
}

func logProduceMessage(deliveryChan chan kafka.Event) {
	kafkaEvent := <- deliveryChan
	kafkaMessage := kafkaEvent.(*kafka.Message)

	if kafkaMessage.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", kafkaMessage.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s - partition [%d] - offset %v\n",
			*kafkaMessage.TopicPartition.Topic, kafkaMessage.TopicPartition.Partition, kafkaMessage.TopicPartition.Offset)
	}
}