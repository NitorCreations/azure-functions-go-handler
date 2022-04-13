// Package handler provides the default custom Go handler logic for
// Go Function Apps.
package handler

import (
	"encoding/json"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

// InvokeRequest models the data sent by the Funtions Runtime for
// the custom Go handler.
type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]json.RawMessage
}

// InvokeRequest models the data expected by the Functions Runtime
// as a response to an invocation.
type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

// Handler runs the custom handler HTTP server and runs the registered
// methods upon invocation requests.
type Handler struct {
	Debug     bool
	Functions map[string]*function.Function
}
