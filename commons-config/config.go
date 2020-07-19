package commons_config

import "os"

const (
	kafkaEnvKey             = "KAFKA_BOOTSTRAP_SERVERS"
	defaultBootstrapServers = "localhost:9092"
)

func KafkaBootstrapServersFromEnvOrDefault() string {
	bootstrapServers := os.Getenv(kafkaEnvKey)
	if bootstrapServers == "" {
		return defaultBootstrapServers
	}

	return bootstrapServers
}
