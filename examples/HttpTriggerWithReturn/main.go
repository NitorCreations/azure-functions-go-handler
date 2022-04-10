package HttpTriggerWithReturn

import (
	"log"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func Handle(ctx *function.Context, req *function.HttpRequest) function.HttpResponse {
	log.Println("Trace logs")				 // you can use this for debugging implementation
	ctx.Log.Println("Function logs") // use this for function logging
	
	return function.HttpResponse{
		Body: function.H{
			"message": "Hello world",
		},
		Headers: function.ContentTypeJson,
	}
}
