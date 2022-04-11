package function

import (
	"encoding/json"
	"strings"
)

func parse[V any](src RawData, path string) (V, bool) {
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
				return res, true
			}
		} else {
			if err := json.Unmarshal(val, &data); err != nil {
				break
			}
		}
	}

	return res, false
}

func (r RawData) Bool(path string) (bool, bool) {
	return parse[bool](r, path)
}

func (r RawData) Int(path string) (int, bool) {
	return parse[int](r, path)
}

func (r RawData) Int8(path string) (int8, bool) {
	return parse[int8](r, path)
}

func (r RawData) Int16(path string) (int16, bool) {
	return parse[int16](r, path)
}

func (r RawData) Int32(path string) (int32, bool) {
	return parse[int32](r, path)
}

func (r RawData) Int64(path string) (int64, bool) {
	return parse[int64](r, path)
}

func (r RawData) Uint(path string) (uint, bool) {
	return parse[uint](r, path)
}

func (r RawData) Uint8(path string) (uint8, bool) {
	return parse[uint8](r, path)
}

func (r RawData) Uint16(path string) (uint16, bool) {
	return parse[uint16](r, path)
}

func (r RawData) Uint32(path string) (uint32, bool) {
	return parse[uint32](r, path)
}

func (r RawData) Uint64(path string) (uint64, bool) {
	return parse[uint64](r, path)
}

func (r RawData) Float(path string) (float32, bool) {
	return parse[float32](r, path)
}

func (r RawData) Float32(path string) (float32, bool) {
	return parse[float32](r, path)
}

func (r RawData) Float64(path string) (float64, bool) {
	return parse[float64](r, path)
}

func (r RawData) String(path string) (string, bool) {
	return parse[string](r, path)
}

func (r RawData) Map(path string) (map[string]interface{}, bool) {
	return parse[map[string]interface{}](r, path)
}

func (r RawData) Slice(path string) ([]interface{}, bool) {
	return parse[[]interface{}](r, path)
}
