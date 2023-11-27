package generic

import "reflect"

// ternary operator
func If[T any](condition bool, resultTrue, resultFalse T) T {
	if condition {
		return resultTrue
	}
	return resultFalse
}

// returns pointer of a value
func Ptr[T any](v T) *T {
	return &v
}

// returns first non-zero/empty value
func Coalesce[T comparable](values ...T) (v T) {
	var zero T
	for _, v = range values {
		if v != zero {
			return
		}
	}
	return
}

func IsEmpty(value interface{}) bool {
	switch t := value.(type) {
	case string:
		return len(t) == 0
	case []interface{}:
		return len(t) == 0
	case map[interface{}]interface{}:
		return len(t) == 0
	case *interface{}:
		return *t == nil
	default:
		// Check for zero values for primitive data types
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			return value == reflect.ValueOf(nil).Interface()
		}
		return value == reflect.ValueOf(0).Interface()
	}
}
