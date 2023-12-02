package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"
)

func Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
