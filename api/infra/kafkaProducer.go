package infra

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelneri/microservices-orchestration/commons/config"
)

func CreateKafkaProducer() *kafka.Producer {
	bootstrapServers := config.KafkaBootstrapServersFromEnvOrDefault()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	log.Print("Kafka Producer started. bootstrapServers: ", bootstrapServers)

	return producer
}

func ProduceMessage(key, value []byte, producer *kafka.Producer, topic string, deliveryChan chan kafka.Event) error {
	produceError := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value,
	}, deliveryChan)

	if produceError != nil {
		log.Panicln("error to produce: ", produceError)
	}

	logProduceMessage(deliveryChan)

	return produceError
}

func logProduceMessage(deliveryChan chan kafka.Event) {
	kafkaEvent := <-deliveryChan
	kafkaMessage := kafkaEvent.(*kafka.Message)

	if kafkaMessage.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", kafkaMessage.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s - partition [%d] - offset %v\n",
			*kafkaMessage.TopicPartition.Topic, kafkaMessage.TopicPartition.Partition, kafkaMessage.TopicPartition.Offset)
	}
}
