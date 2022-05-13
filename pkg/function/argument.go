package function

import (
	"encoding/json"
	"reflect"
)

func NewArgument(name string, argType reflect.Type, direction Direction) (*Argument, error) {
	indirect := true

	if argType.Kind() == reflect.Ptr {
		indirect = false
		argType = argType.Elem()
	}

	return &Argument{
		Name:      name,
		Type:      argType,
		Direction: direction,
		indirect:  indirect,
	}, nil
}

func (a *Argument) Instance() *Argument {
	return &Argument{
		Name:      a.Name,
		Type:      a.Type,
		Value:     reflect.New(a.Type),
		Direction: a.Direction,
		indirect:  a.indirect,
	}
}

func (a *Argument) Allocate() error {
	switch t := a.Value.Interface().(type) {
	case Allocatable:
		if err := t.Allocate(); err != nil {
			return err
		}
	}
	return nil
}

func (a *Argument) Read(in map[string]json.RawMessage) error {
	// make sure data is allocated
	if err := a.Allocate(); err != nil {
		return err
	}

	if data, ok := in[a.Name]; ok {
		return json.Unmarshal(data, a.Value.Interface())
	}
	return ErrBindingNotFound
}

func (a *Argument) Write(out map[string]any) error {
	out[a.Name] = a.Argument().Interface()
	return nil
}

func (a *Argument) Argument() reflect.Value {
	if a.indirect {
		return reflect.Indirect(a.Value)
	}
	return a.Value
}
