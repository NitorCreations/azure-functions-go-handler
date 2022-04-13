package function

import "reflect"

func validCtxArg(argType reflect.Type) bool {
	if argType.Kind() != reflect.Ptr {
		return false
	}
	return argType.Elem().Name() == "Context"
}

func NewFunction(config *Config, reference any) (*Function, error) {
	funType := reflect.TypeOf(reference)
	if funType.Kind() != reflect.Func {
		return nil, ErrBadReference
	}

	arguments := make([]Argument, funType.NumIn())

	for i := 0; i < funType.NumIn(); i++ {
		argType := funType.In(i)

		if i == 0 {
			if !validCtxArg(argType) {
				return nil, ErrBadReference
			}
			arguments[i] = *NewArgument(
				"_ctx", argType, DirectionIn)

		} else {
			binding := config.Bindings[i-1]
			arguments[i] = *NewArgument(
				binding.Name, argType, binding.Direction)
		}
	}

	return &Function{
		Config:    *config,
		Reference: reference,
		Arguments: arguments,
	}, nil
}
