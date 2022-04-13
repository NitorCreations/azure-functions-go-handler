package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

// Start the custom handler HTTP server. Default serve port is 8080,
// overridable with environment variable FUNCTIONS_CUSTOMHANDLER_PORT.
func (w *Handler) Start() error {
	port := "8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = val
	}

	http.HandleFunc("/", w.handle)
	log.Printf("Handler process listening in 127.0.0.1:%s", port)
	return http.ListenAndServe(":"+port, nil)
}

func (w *Handler) handle(res http.ResponseWriter, req *http.Request) {
	// Debug request
	if w.Debug {
		dump, _ := httputil.DumpRequest(req, true)
		log.Printf("REQUEST:\n%s", string(dump))
	}

	// Defer panic handling
	defer func() {
		if r := recover(); r != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Printf("%s", r)
		}
	}()

	if req.Method == http.MethodPost {
		// Parse invoke request
		request := parseInvokeRequest(req)
		sys, err := request.sys()
		panicIf(err, "Failed to parse invoke request")

		// Resolve method
		fun, ok := w.Functions[sys.MethodName]
		if !ok {
			panicMsg("No handler found for function %s", sys.MethodName)
		}

		// Invoke
		response := invoke(request, fun)
		response.encode(res)
	}
}

func panicIf(err error, format string, a ...any) {
	if err != nil {
		if format != "" {
			panic(fmt.Sprintf(format, a...))
		} else {
			panic(err)
		}
	}
}

func panicMsg(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
