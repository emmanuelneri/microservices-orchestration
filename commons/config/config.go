package config

import "os"

const (
	kafkaBootsrapServersKey   = "KAFKA_BOOTSTRAP_SERVERS"
	defaultBootstrapServers   = "localhost:9092"
	kafkaSchemaRegistryUrlKey = "KAFKA_SCHEMA_REGISTRY_URL"
	defaultSchemaRegistryUrl  = "localhost:8081"
)

func KafkaBootstrapServersFromEnvOrDefault() string {
	return getEnvOrDefault(kafkaBootsrapServersKey, defaultBootstrapServers)
}

func SchemaRegistryUrlFromEnvOrDefault() string {
	return getEnvOrDefault(kafkaSchemaRegistryUrlKey, defaultSchemaRegistryUrl)
}

func getEnvOrDefault(envKey, defaultValue string) string {
	envValue := os.Getenv(envKey)
	if envValue == "" {
		return defaultValue
	}

	return envKey
}
