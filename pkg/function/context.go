package function

import "fmt"

func (l *Log) Print(a ...any) {
	l.ctx.Logs = append(l.ctx.Logs, fmt.Sprint(a...))
}

func (l *Log) Printf(format string, a ...any) {
	l.ctx.Logs = append(l.ctx.Logs, fmt.Sprintf(format, a...))
}

func (l *Log) Println(a ...any) {
	l.ctx.Logs = append(l.ctx.Logs, fmt.Sprintln(a...))
}

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
