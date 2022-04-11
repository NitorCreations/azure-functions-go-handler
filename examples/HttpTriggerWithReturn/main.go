package HttpTriggerWithReturn

import (
	"log"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func Handle(ctx *function.Context, req *function.HttpRequest) function.HttpResponse {
	log.Println("Trace logs")        // you can use this for debugging implementation
	ctx.Log.Println("Function logs") // use this for function logging

	// Parse JSON body as map[string]any
	if req.Body != "" && req.IsJSON() {
		data := make(function.H)
		if err := req.BodyJSON(&data); err != nil {
			return function.HttpResponse{
				Status: 400,
			}
		}
		ctx.Log.Println("JSON body as map:", data)
	}

	// Access raw invocation data
	ctx.Log.Println(ctx.Metadata.String("sys.UtcNow"))

	return function.HttpResponse{
		Body: function.H{
			"message": "Hello world",
		},
		Headers: function.ContentTypeJson,
	}
}
