package function

import "fmt"

// Print appends to function execution context log.
// Arguments are handled in the manner of fmt.Sprint.
func (l *Log) Print(a ...any) {
	l.ctx.Logs = append(l.ctx.Logs, fmt.Sprint(a...))
}

// Printf appends to function execution context log.
// Arguments are handled in the manner of fmt.Sprintf.
func (l *Log) Printf(format string, a ...any) {
	l.ctx.Logs = append(l.ctx.Logs, fmt.Sprintf(format, a...))
}

// Println appends to function execution context log.
// Arguments are handled in the manner of fmt.Sprintln.
func (l *Log) Println(a ...any) {
	l.ctx.Logs = append(l.ctx.Logs, fmt.Sprintln(a...))
}

// NewContext creates a new context struct from provided
// invocation data.
func NewContext(data RawData, metadata RawData) *Context {
	ctx := &Context{
		Data:     data,
		Metadata: metadata,
		Outputs:  make(map[string]interface{}),
		Logs:     []string{},
	}
	ctx.Log = Log{ctx: ctx}
	return ctx
}
