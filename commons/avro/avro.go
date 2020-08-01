package avro

import (
	"github.com/linkedin/goavro/v2"
	"io/ioutil"
)

func LoadAvroCodec(filename string) (*goavro.Codec, error) {
	schema, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return goavro.NewCodec(string(schema))
}
