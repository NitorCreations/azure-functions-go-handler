package handler

import (
	"encoding/json"
)

type System struct {
	MethodName string
	UtcNow string
	RandGuid string
}

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]json.RawMessage
}

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

type Handler struct {
	Debug bool
	Methods map[string]interface{}
}
