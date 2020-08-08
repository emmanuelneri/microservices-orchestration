package kafka

import (
	"io/ioutil"

	"github.com/riferrei/srclient"
)

const schemaType = "AVRO"

func getOrCreateSchema(schemaRegistryUrl, topic, filename string) (*srclient.Schema, error) {
	schemaRegistryClient := srclient.CreateSchemaRegistryClient(schemaRegistryUrl)
	schema, err := geSchema(schemaRegistryUrl, topic)
	if schema == nil {
		schemaBytes, _ := ioutil.ReadFile(filename)
		return schemaRegistryClient.CreateSchema(topic, string(schemaBytes), schemaType, false)
	}

	return schema, err
}

func geSchema(schemaRegistryUrl, topic string) (*srclient.Schema, error) {
	schemaRegistryClient := srclient.CreateSchemaRegistryClient(schemaRegistryUrl)
	return schemaRegistryClient.GetLatestSchema(topic, false)
}
