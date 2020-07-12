package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func SubscribeConsumer(bootstrapServers string, consumerGroupName string, topic string) {
	consumer, consumerError := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          consumerGroupName,
		"auto.offset.reset": "earliest",
	})

	if consumerError != nil {
		log.Fatal("Failed to create consumer.", consumerError)
	}

	subscriberError := consumer.Subscribe(topic, nil)

	if subscriberError != nil {
		log.Fatal("Failed to subscribe topic: ", topic, subscriberError)
	}

	for true {
		eventConsumed := consumer.Poll(0)
		switch event := eventConsumed.(type) {
		case *kafka.Message:
			fmt.Printf("%% Message Consumed. partition: [%d] - key: %s - value: %s\n", event.TopicPartition.Partition, string(event.Key), string(event.Value))
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", event)
		case kafka.Error:
			log.Panicln("Message error Error: %v\n", event)
		default:
		}
	}
}
