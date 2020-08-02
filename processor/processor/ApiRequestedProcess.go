package processor

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaSubscriber "github.com/emmanuelneri/microservices-orchestration/commons/kafka"
	"github.com/linkedin/goavro/v2"
)

const (
	topic = "ApiRequested"
)

func ApiRequestedSubscriber(consumer *kafka.Consumer, codec *goavro.Codec) {
	subscriber := kafkaSubscriber.CreateSubscriber(consumer, topic, codec)
	subscriber.Subscribe()

	go func() {
		for {
			message := <-subscriber.ConsumeChan
			fmt.Printf("%% Message Consumed. headers: %s - partition: [%d] - key: %s - value: %s\n",
				message.Headers, message.Partition, string(message.Key), message.ValueAsTextual)
		}
	}()

}
