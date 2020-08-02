package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelneri/microservices-orchestration/commons/config"
	"github.com/linkedin/goavro/v2"
)

// Producer: encapsulate kafka producer and configurations to produce message
type Producer struct {
	kafkaProducer *kafka.Producer
	topic         string
	codec         *goavro.Codec
	deliveryChan  chan kafka.Event
}

func CreateProducer(topic string, codec *goavro.Codec) *Producer {
	return &Producer{kafkaProducer: createKafkaProducer(), topic: topic, codec: codec, deliveryChan: make(chan kafka.Event, 10000)}
}

func createKafkaProducer() *kafka.Producer {
	bootstrapServers := config.KafkaBootstrapServersFromEnvOrDefault()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	log.Print("Kafka Producer started. bootstrapServers: ", bootstrapServers)

	return producer
}

func (producer *Producer) Produce(key []byte, avroMessage AvroMessage) error {
	binary, err := serialize(avroMessage, producer.codec)
	if err != nil {
		log.Panicln("serialize error: ", err)
	}

	err = producer.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &producer.topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          binary,
	}, producer.deliveryChan)

	if err != nil {
		log.Panicln("produce error: ", err)
	}

	logProduceMessage(producer.deliveryChan)

	return err
}

func serialize(avroMessage AvroMessage, codec *goavro.Codec) ([]byte, error) {
	return codec.BinaryFromNative(nil, avroMessage.ToMap())
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
