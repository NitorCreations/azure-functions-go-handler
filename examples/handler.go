// Code generated by gofunc. Regenarate with cmd: gofunc generate
package main

import (
	"net/http"

	HttpTrigger "github.com/NitorCreations/azure-functions-go-handler/examples/HttpTrigger"
	HttpTriggerWithReturn "github.com/NitorCreations/azure-functions-go-handler/examples/HttpTriggerWithReturn"
	"github.com/NitorCreations/azure-functions-go-handler/pkg/handler"
)

func main() {
	handler := &handler.Handler{
		Debug: false,
		Methods: map[string]interface{}{
			"HttpTrigger":           HttpTrigger.Handle,
			"HttpTriggerWithReturn": HttpTriggerWithReturn.Handle,
		},
	}

	err := handler.Start()
	if err != http.ErrServerClosed {
		panic(err)
	}
}
