package subscriber

import (
	"fmt"
	avro "github.com/linkedin/goavro/v2"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	topic = "ApiRequested"
)

func ApiRequestedSubscriber(consumer *kafka.Consumer, codec *avro.Codec) {
	subscriberError := consumer.Subscribe(topic, nil)

	if subscriberError != nil {
		log.Fatal("Failed to subscribe topic: ", topic, subscriberError)
	}

	for true {
		eventConsumed := consumer.Poll(0)
		switch event := eventConsumed.(type) {
		case *kafka.Message:
			native, _, _ := codec.NativeFromBinary(event.Value)
			textual, _  := codec.TextualFromNative(nil, native)
			fmt.Printf("%% Message Consumed. headers: %s - partition: [%d] - key: %s - value: %s\n",
				event.Headers, event.TopicPartition.Partition, string(event.Key), textual)
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", event)
		case kafka.Error:
			log.Panicln("Message error Error: %v\n", event)
		default:
		}
	}
}
