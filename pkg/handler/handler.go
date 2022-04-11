package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

// Invocation system metadata
type system struct {
	MethodName string
	UtcNow     string
	RandGuid   string
}

// Start the custom handler HTTP server. Default serve port is 8080,
// overridable with environment variable FUNCTIONS_CUSTOMHANDLER_PORT.
func (w *Handler) Start() error {
	port := "8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = val
	}

	http.HandleFunc("/", w.handle)
	log.Printf("Handler process listening in 127.0.0.1:%s", port)
	return http.ListenAndServe(":"+port, nil)
}

func (w *Handler) handle(res http.ResponseWriter, req *http.Request) {
	// Debug request
	if w.Debug {
		dump, _ := httputil.DumpRequest(req, true)
		log.Printf("REQUEST:\n%s", string(dump))
	}

	// Defer panic handling
	defer func() {
		if r := recover(); r != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Printf("%s", r)
		}
	}()

	if req.Method == http.MethodPost {
		// Parse invoke request
		invokeRequest := parseInvokeRequest(req)
		sys, err := invokeRequest.sys()
		panicIf(err, "Failed to parse invoke request")

		// Resolve method
		method, ok := w.Methods[sys.MethodName]
		if !ok {
			panicMsg("No handler found for method %s", sys.MethodName)
		}

		methodType := reflect.TypeOf(method)
		if methodType.Kind() != reflect.Func {
			panicMsg("Method %s not a function", sys.MethodName)
		}

		// Handler method inputs
		context := function.NewContext(invokeRequest.Data, invokeRequest.Metadata)
		in := []reflect.Value{reflect.ValueOf(context)}

		for i := 1; i < methodType.NumIn(); i++ {
			inputName := reflect.ValueOf(invokeRequest.Data).MapKeys()[i-1].String()
			if rv, ok := invokeRequest.Data[inputName]; ok {
				indirect := true
				inputType := methodType.In(i)

				if inputType.Kind() == reflect.Pointer {
					indirect = false
					inputType = inputType.Elem()
				}

				value := reflect.New(inputType)
				err := json.Unmarshal(rv, value.Interface())
				panicIf(err, "Unable to parse input parameter %s", inputName)

				if indirect {
					in = append(in, reflect.Indirect(value))
				} else {
					in = append(in, value)
				}
			}
		}

		// Invoke method
		out := reflect.ValueOf(method).Call(in)

		// Handle method outputs
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

		invokeResponse.encode(res)
	}
}

func panicIf(err error, format string, a ...any) {
	if err != nil {
		if format != "" {
			panic(fmt.Sprintf(format, a...))
		} else {
			panic(err)
		}
	}
}

func panicMsg(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
