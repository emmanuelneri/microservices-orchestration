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

func CreateProducer(topic, schemaFile string) (*Producer, error) {
	schemaRegistryUrl := config.SchemaRegistryUrlFromEnvOrDefault()
	schema, err := getOrCreateSchema(schemaRegistryUrl, topic, schemaFile)
	if err != nil {
		return nil, err
	}

	producer, err := createKafkaProducer()
	if err != nil {
		return nil, err
	}

	return &Producer{kafkaProducer: producer,
			topic:        topic,
			codec:        schema.Codec(),
			deliveryChan: make(chan kafka.Event)},
		err
}

func createKafkaProducer() (*kafka.Producer, error) {
	bootstrapServers := config.KafkaBootstrapServersFromEnvOrDefault()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, err
	}

	log.Print("Kafka Producer started. bootstrapServers: ", bootstrapServers)

	return producer, nil
}

func (producer *Producer) Produce(key []byte, value map[string]interface{}) error {
	binary, err := serialize(value, producer.codec)
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

func serialize(value map[string]interface{}, codec *goavro.Codec) ([]byte, error) {
	return codec.BinaryFromNative(nil, value)
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
