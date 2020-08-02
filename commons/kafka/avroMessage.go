package kafka

type AvroMessage interface {
	ToMap() map[string]interface{}
}
