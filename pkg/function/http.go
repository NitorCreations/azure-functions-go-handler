package function

import (
	"encoding/json"
)

type HttpRequestHeaders map[string][]string

type HttpRequest struct {
	Url     string
	Method  string
	Query   map[string]string
	Headers HttpRequestHeaders
	Params  map[string]string
	// Identities []struct{
	// 	AuthenticationType string
	// 	IsAuthenticated    bool
	// 	Actor              string
	// 	BootstrapContext   string
	// 	Claims             []string
	// 	Label              string
	// 	Name               string
	// 	NameClaimType      string
	// 	RoleClaimType      string
	// }
}

type HttpResponseHeaders map[string][]string

func (h *HttpResponseHeaders) ContentType(val string) *HttpResponseHeaders {
	(*h)["Content-Type"] = []string{val}
	return h
}

func (h *HttpResponseHeaders) ContentTypeJson() *HttpResponseHeaders {
	return h.ContentType("application/json")
}

func (h *HttpResponseHeaders) ContentTypeText() *HttpResponseHeaders {
	return h.ContentType("text/plain")
}

var (
	ContentTypeJson HttpResponseHeaders = *(&HttpResponseHeaders{}).ContentTypeJson()
	ContentTypeText HttpResponseHeaders = *(&HttpResponseHeaders{}).ContentTypeText()
)

type HttpResponse struct {
	Status  uint32
	Body    interface{}
	Headers HttpResponseHeaders
}

func (h HttpResponse) MarshalJSON() ([]byte, error) {
	if h.Status == 0 {
		h.Status = 200
	}

	return json.Marshal(&struct {
		Status  uint32    				  `json:"status"`
		Body    interface{} 			  `json:"body"`
		Headers HttpResponseHeaders `json:"headers"`
	}{
		h.Status, h.Body, h.Headers,
	})
}
