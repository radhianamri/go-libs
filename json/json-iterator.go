package json

import (
	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.Config{
	EscapeHTML:             false,
	UseNumber:              true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

func Unmarshal(body []byte, v interface{}) error {
	return Json.Unmarshal(body, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return Json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	return Json.MarshalToString(v)
}
