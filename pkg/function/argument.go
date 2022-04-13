package function

import (
	"encoding/json"
	"reflect"
)

func NewArgument(name string, argType reflect.Type, direction Direction) *Argument {
	indirect := true

	if argType.Kind() == reflect.Pointer {
		indirect = false
		argType = argType.Elem()
	}

	value := reflect.New(argType)

	return &Argument{
		Name:      name,
		Value:     value,
		Direction: direction,
		indirect:  indirect,
	}
}

func (b *Argument) Read(in map[string]json.RawMessage) error {
	if data, ok := in[b.Name]; ok {
		return json.Unmarshal(data, b.Value.Interface())
	}
	return ErrBindingNotFound
}

func (b *Argument) Write(out map[string]any) error {
	out[b.Name] = b.Value.Interface()
	return nil
}

func (b *Argument) Argument() reflect.Value {
	if b.indirect {
		return reflect.Indirect(b.Value)
	}
	return b.Value
}
