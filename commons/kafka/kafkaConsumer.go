package kafka

import (
	"fmt"
	"github.com/emmanuelneri/microservices-orchestration/commons/config"
	"github.com/linkedin/goavro/v2"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Subscriber: encapsulate topic subscription
type Subscriber struct {
	kafkaConsumer *kafka.Consumer
	topic         string
	codec         *goavro.Codec
	ConsumeChan   chan *ConsumedMessage
}

func CreateSubscriber(kafkaConsumer *kafka.Consumer, topic string) (*Subscriber, error) {
	schemaRegistryUrl := config.SchemaRegistryUrlFromEnvOrDefault()
	schema, err := geSchema(schemaRegistryUrl, topic)
	if err != nil {
		return nil, err
	}

	return &Subscriber{kafkaConsumer: kafkaConsumer,
			topic:       topic,
			codec:       schema.Codec(),
			ConsumeChan: make(chan *ConsumedMessage)},
		err
}

type ConsumedMessage struct {
	Key            []byte
	Value          []byte
	ValueAsNative  interface{}
	ValueAsTextual []byte
	Partition      int32
	Timestamp      time.Time
	Headers        map[string][]byte
}

func createConsumedMessage(message *kafka.Message, native interface{}, textual []byte) *ConsumedMessage {
	headers := make(map[string][]byte)
	for _, header := range message.Headers {
		headers[header.Key] = header.Value
	}
	return &ConsumedMessage{
		Key:            message.Key,
		Value:          message.Value,
		ValueAsNative:  native,
		ValueAsTextual: textual,
		Partition:      message.TopicPartition.Partition,
		Timestamp:      message.Timestamp,
		Headers:        headers,
	}
}

func CreateConsumer(consumerGroupName string) *kafka.Consumer {
	bootstrapServers := config.KafkaBootstrapServersFromEnvOrDefault()

	consumer, consumerError := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          consumerGroupName,
		"auto.offset.reset": "earliest",
	})

	if consumerError != nil {
		log.Fatal("Failed to create consumer.", consumerError)
	}

	log.Print("Kafka Consumer started. bootstrapServers: ", bootstrapServers)

	return consumer
}

func (subscribe *Subscriber) Subscribe() {
	subscriberError := subscribe.kafkaConsumer.Subscribe(subscribe.topic, nil)

	if subscriberError != nil {
		log.Fatal("Failed to subscribe topic: ", subscribe.topic, subscriberError)
	}

	go func() {
		for true {
			eventConsumed := subscribe.kafkaConsumer.Poll(0)
			switch event := eventConsumed.(type) {
			case *kafka.Message:
				native, _, err := subscribe.codec.NativeFromBinary(event.Value)
				textual, err := subscribe.codec.TextualFromNative(nil, native)

				if err != nil {
					log.Panicf("Message deserialize error: %v\n", err)
				} else {
					subscribe.ConsumeChan <- createConsumedMessage(event, native, textual)
				}
			case kafka.PartitionEOF:
				fmt.Printf("PartitionEOF %v\n", event)
			case kafka.Error:
				log.Panicf("Message error Error: %v\n", event)
			default:
			}
		}
	}()
}
