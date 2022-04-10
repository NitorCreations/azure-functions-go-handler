package handler

import (
	"encoding/json"
	"net/http"
)

func parseInvokeRequest(req *http.Request) *InvokeRequest {
	var invokeRequest InvokeRequest
	d := json.NewDecoder(req.Body)
	d.Decode(&invokeRequest)
	return &invokeRequest
}

func (i *InvokeRequest) sys() (*System, error) {
	var sys System
	err := json.Unmarshal(i.Metadata["sys"], &sys)
	return &sys, err
}

func (i *InvokeResponse) encode(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(i)
	panicIf(err, "Failed to encode invoke response %s", err)
}
