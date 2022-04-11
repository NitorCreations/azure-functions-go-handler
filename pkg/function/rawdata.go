package function

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrUnknownProperty = errors.New("unknown property")
)

func parse[V any](src RawData, path string) (V, error) {
	var res V
	data := src

	keys := strings.Split(path, ".")
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val, ok := data[key]

		if !ok {
			break
		}

		if i+1 == len(keys) {
			if err := json.Unmarshal(val, &res); err == nil {
				return res, nil
			}
		} else {
			if err := json.Unmarshal(val, &data); err != nil {
				break
			}
		}
	}

	return res, ErrUnknownProperty
}

// Bool unmarshals a bool value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Bool(path string) (bool, error) {
	return parse[bool](r, path)
}

// Int unmarshals an int value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Int(path string) (int, error) {
	return parse[int](r, path)
}

// Int8 unmarshals an int8 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Int8(path string) (int8, error) {
	return parse[int8](r, path)
}

// Int16 unmarshals an int16 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Int16(path string) (int16, error) {
	return parse[int16](r, path)
}

// Int32 unmarshals an int32 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Int32(path string) (int32, error) {
	return parse[int32](r, path)
}

// Int64 unmarshals an int64 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Int64(path string) (int64, error) {
	return parse[int64](r, path)
}

// Uint unmarshals an uint value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Uint(path string) (uint, error) {
	return parse[uint](r, path)
}

// Uint8 unmarshals an uint8 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Uint8(path string) (uint8, error) {
	return parse[uint8](r, path)
}

// Uint16 unmarshals an uint16 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Uint16(path string) (uint16, error) {
	return parse[uint16](r, path)
}

// Uint32 unmarshals an uint32 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Uint32(path string) (uint32, error) {
	return parse[uint32](r, path)
}

// Uint64 unmarshals an uint64 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Uint64(path string) (uint64, error) {
	return parse[uint64](r, path)
}

// Float unmarshals a float32 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Float(path string) (float32, error) {
	return parse[float32](r, path)
}

// Float32 unmarshals a float32 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Float32(path string) (float32, error) {
	return parse[float32](r, path)
}

// Float64 unmarshals a float64 value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Float64(path string) (float64, error) {
	return parse[float64](r, path)
}

// String unmarshals a string value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) String(path string) (string, error) {
	return parse[string](r, path)
}

// Map unmarshals a map[string]interface{} value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Map(path string) (map[string]interface{}, error) {
	return parse[map[string]interface{}](r, path)
}

// Slice unmarshals a []]interface{} value from the given path.
// The path is a property path with support for dot notation.
func (r RawData) Slice(path string) ([]interface{}, error) {
	return parse[[]interface{}](r, path)
}
