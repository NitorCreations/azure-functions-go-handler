package main

import (
	"net/http"

	"github.com/NitorCreations/azure-functions-go-handler/examples/HttpTrigger"
	"github.com/NitorCreations/azure-functions-go-handler/examples/HttpTriggerWithReturn"
	worker "github.com/NitorCreations/azure-functions-go-handler/pkg/handler"
)

func main() {
	worker := &worker.Handler{
		Debug: true,
		Methods: map[string]interface{}{
			"HttpTrigger": HttpTrigger.Handle,
			"HttpTriggerWithReturn": HttpTriggerWithReturn.Handle,
		},
	}

	err := worker.Start()
	if err != http.ErrServerClosed {
		panic(err)
	}
}
