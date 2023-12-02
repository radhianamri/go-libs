package json

import "github.com/valyala/fastjson"

func GetBool(data []byte, keys ...string) bool {
	return fastjson.GetBool(data, keys...)
}

func GetFloat64(data []byte, keys ...string) float64 {
	return fastjson.GetFloat64(data, keys...)
}

func GetInt(data []byte, keys ...string) int {
	return fastjson.GetInt(data, keys...)
}

func GetString(data []byte, keys ...string) string {
	return fastjson.GetString(data, keys...)
}

func Validate(s string) error {
	return fastjson.Validate(s)
}

func ValidateBytes(b []byte) error {
	return fastjson.ValidateBytes(b)
}

func Exists(data []byte, keys ...string) bool {
	return fastjson.Exists(data, keys...)
}

func Parse(s string) (*fastjson.Value, error) {
	return fastjson.Parse(s)
}

func ParseBytes(b []byte) (*fastjson.Value, error) {
	return fastjson.ParseBytes(b)
}
