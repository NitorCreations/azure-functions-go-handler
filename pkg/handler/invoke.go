package handler

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

// Invocation system metadata
type system struct {
	MethodName string
	UtcNow     string
	RandGuid   string
}

func parseInvokeRequest(req *http.Request) *InvokeRequest {
	var invokeRequest InvokeRequest
	d := json.NewDecoder(req.Body)
	err := d.Decode(&invokeRequest)
	panicIf(err, "Failed to parse invode request")
	return &invokeRequest
}

func (i *InvokeRequest) sys() (*system, error) {
	var sys system
	err := json.Unmarshal(i.Metadata["sys"], &sys)
	return &sys, err
}

func invoke(request *InvokeRequest, fun *function.Function) *InvokeResponse {
	// Invocation context
	context := function.NewContext(request.Data, request.Metadata)

	// Handler method inputs
	in := make([]reflect.Value, len(fun.Arguments))
	for i, arg := range fun.Arguments {
		if i == 0 {
			in[i] = reflect.ValueOf(context)
			continue
		}

		if arg.Direction != function.DirectionOut {
			panicIf(arg.Read(request.Data),
				"Failed to parse binding %s", arg.Name)
		}
		in[i] = arg.Argument()
	}

	// Invoke method
	out := reflect.ValueOf(fun.Reference).Call(in)

	// Handle function outputs
	for _, arg := range fun.Arguments {
		if arg.Direction != function.DirectionIn {
			arg.Write(context.Outputs)
		}
	}

	// Handle function return value
	var returnValue interface{} = nil
	if len(out) > 0 {
		v := out[0]
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		returnValue = v.Interface()
	}

	// Build invoke response
	invokeResponse := InvokeResponse{
		Logs:        context.Logs,
		Outputs:     context.Outputs,
		ReturnValue: returnValue,
	}

	return &invokeResponse
}

func (i *InvokeResponse) encode(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(i)
	panicIf(err, "Failed to encode invoke response %s", err)
}
