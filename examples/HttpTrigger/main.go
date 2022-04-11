package HttpTrigger

import (
	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func Handle(ctx *function.Context) {
	ctx.Outputs["res"] = function.H{
		"status": 200,
		"body":   "Hello Gophers!",
	}
}
