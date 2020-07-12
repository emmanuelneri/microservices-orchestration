package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	bootstrapServers := getBootstrapServers()
	log.Print("API started")
	log.Print("Kafka bootstrapServers: ", bootstrapServers)

	topic := "ApiRequested"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		requestBody, bodyError := ioutil.ReadAll(request.Body)
		if bodyError != nil {
			panic(bodyError)
		}

		log.Println("API requested: ", string(requestBody))

		deliveryChan := make(chan kafka.Event, 10000)
		produceError := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic},
			Value:          requestBody,
		}, deliveryChan)

		kafkaEvent := <- deliveryChan
		kafkaMessage := kafkaEvent.(*kafka.Message)

		if kafkaMessage.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", kafkaMessage.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s - partition [%d] - offset %v\n",
				*kafkaMessage.TopicPartition.Topic, kafkaMessage.TopicPartition.Partition, kafkaMessage.TopicPartition.Offset)
		}

		if produceError != nil {
			panic(produceError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO Configmap
func getBootstrapServers() string {
	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if bootstrapServers == "" {
		return "localhost:9092"
	}

	return bootstrapServers
}
