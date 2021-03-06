package config

import "os"

const (
	kafkaBootsrapServersKey   = "KAFKA_BOOTSTRAP_SERVERS"
	defaultBootstrapServers   = "localhost:9092"
	kafkaSchemaRegistryUrlKey = "KAFKA_SCHEMA_REGISTRY_URL"
	defaultSchemaRegistryUrl  = "http://localhost:8081"
	syncProcessorAppKey       = "SYNC_PROCESSOR_URL"
	syncProcessorAppUrl       = "http://localhost:9090"
)

func KafkaBootstrapServersFromEnvOrDefault() string {
	return getEnvOrDefault(kafkaBootsrapServersKey, defaultBootstrapServers)
}

func SchemaRegistryUrlFromEnvOrDefault() string {
	return getEnvOrDefault(kafkaSchemaRegistryUrlKey, defaultSchemaRegistryUrl)
}

func SyncProcessorAppUrlFromEnvOrDefault() string {
	return getEnvOrDefault(syncProcessorAppKey, syncProcessorAppUrl)
}

func getEnvOrDefault(envKey, defaultValue string) string {
	envValue := os.Getenv(envKey)
	if envValue == "" {
		return defaultValue
	}

	return envValue
}
