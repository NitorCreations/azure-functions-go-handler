package function

import "encoding/json"

type Log struct {
	ctx *Context
}

type RawData map[string]json.RawMessage

type Context struct {
	Data RawData
	Metadata RawData
	Outputs map[string]interface{}
	Logs []string
	Log Log
}

type H map[string]interface{}
