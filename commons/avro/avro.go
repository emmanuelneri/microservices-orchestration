package avro

import (
	"io/ioutil"

	"github.com/linkedin/goavro/v2"
)

func LoadAvroCodec(filename string) (*goavro.Codec, error) {
	schema, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return goavro.NewCodec(string(schema))
}
