// Package function provides primivites for handler functions.
package function

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Direction string

const (
	DirectionIn    Direction = "in"
	DirectionOut   Direction = "out"
	DirectionInOut Direction = "inout"
)

type DataType string

const (
	DataTypeString DataType = "string"
	DataTypeBinary DataType = "binary"
	DataTypeStream DataType = "stream"
)

type BindingType string

const (
	TypeServiceBusTrigger BindingType = "serviceBusTrigger"
	TypeServiceBus        BindingType = "serviceBus"
	TypeBlobTrigger       BindingType = "blobTrigger"
	TypeBlob              BindingType = "blob"
	TypeManualTrigger     BindingType = "manualTrigger"
	TypeEventHubTrigger   BindingType = "eventHubTrigger"
	TypeEventHub          BindingType = "eventHub"
	TypeTimerTrigger      BindingType = "timerTrigger"
	TypeQueueTrigger      BindingType = "queueTrigger"
	TypeQueue             BindingType = "queue"
	TypeHttpTrigger       BindingType = "httpTrigger"
	TypeHttp              BindingType = "http"
	TypeMobileTable       BindingType = "mobileTable"
	TypeDocumentDB        BindingType = "documentDB"
	TypeTable             BindingType = "table"
	TypeNotificationHub   BindingType = "notificationHub"
	TypeTwilioSms         BindingType = "twilioSms"
	TypeSendGrid          BindingType = "sendGrid"
)

var (
	ErrBadReference    = errors.New("bad reference")
	ErrBindingNotFound = errors.New("binding not found")
)

type Config struct {
	Bindings   []Binding `json:"bindings"`
	Excluded   bool      `json:"excluded,omitempty"`
	EntryPoint string    `json:"entryPoint,omitempty"`
	ScriptFile string    `json:"scriptFile,omitempty"`
}

type Binding struct {
	Name      string      `json:"name"`
	Type      BindingType `json:"type"`
	Direction Direction   `json:"direction"`
	DataType  DataType    `json:"dataType"`
}

type Argument struct {
	Name      string
	Value     reflect.Value
	Direction Direction
	indirect  bool
}

type Function struct {
	Config    Config
	Reference any
	Arguments []Argument
}

// Log provides a scoped access for context logging functions.
type Log struct {
	ctx *Context
}

// RawData type to access data provided in invocation requests.
type RawData map[string]json.RawMessage

// Context for handler functions provides access to invocation data,
// output bindings and function logging.
type Context struct {
	Data     RawData
	Metadata RawData
	Outputs  map[string]interface{}
	Logs     []string
	Log      Log
}

// Shortcut for map[string]interface{}, usefull for defining
// arbitrary JSON structures etc. Shamelessly copied from gin.
type H map[string]interface{}
