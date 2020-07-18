package config

import (
	"os"
	"testing"
)

func TestDefaultKafkaBootstrapServers(test *testing.T) {
	expectedBootstrapServers := "localhost:9092"
	bootstrapServers := KafkaBootstrapServersFromEnvOrDefault()
	if bootstrapServers != expectedBootstrapServers {
		test.Errorf("default bootstrapServers expected %s but was %s",
			expectedBootstrapServers, bootstrapServers)
	}
}

func TestEnvKafkaBootstrapServers(test *testing.T) {
	expectedBootstrapServers := "customHost:9093"
	os.Setenv("KAFKA_BOOTSTRAP_SERVERS",  expectedBootstrapServers)
	bootstrapServers := KafkaBootstrapServersFromEnvOrDefault()
	if bootstrapServers != expectedBootstrapServers {
		test.Errorf("env bootstrapServers expected %s but was %s",
			expectedBootstrapServers, bootstrapServers)
	}
}
