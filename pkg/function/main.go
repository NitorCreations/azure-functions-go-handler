// Package function provides primivites for handler functions.
package function

import "encoding/json"

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
