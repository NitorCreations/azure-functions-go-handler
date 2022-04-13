package function

import (
	"encoding/json"
	"strings"
)

// Shotcut for map[string][]string
type HttpRequestHeaders map[string][]string

// HttpRequest models the HttpTrigger request object.
type HttpRequest struct {
	Url        string
	Method     string
	Body       string
	Query      map[string]string
	Headers    HttpRequestHeaders
	Params     map[string]string
	Identities []struct {
		AuthenticationType string
		IsAuthenticated    bool
		Actor              string
		BootstrapContext   string
		Claims             []string
		Label              string
		Name               string
		NameClaimType      string
		RoleClaimType      string
	}
}

// Parse JSON encoded request body data.
// Argument handled in the manner of json.Unmarshal.
func (h *HttpRequest) BodyJSON(v any) error {
	return json.Unmarshal([]byte(h.Body), v)
}

// IsJson checks if the request body is JSON encoded.
func (h *HttpRequest) IsJSON() bool {
	if val, ok := h.Headers["Content-Type"]; ok {
		for _, ct := range val {
			if strings.HasPrefix(ct, "application/json") {
				return true
			}
		}
	}
	return false
}

// Shotcut for map[string][]string
type HttpResponseHeaders map[string][]string

// ContentType sets the HttpResponseHeaders content type.
func (h *HttpResponseHeaders) ContentType(val string) *HttpResponseHeaders {
	(*h)["Content-Type"] = []string{val}
	return h
}

// Shortcut for ContentType("application/json")
func (h *HttpResponseHeaders) ContentTypeJson() *HttpResponseHeaders {
	return h.ContentType("application/json")
}

// Shortcut for ContentType("text/plain")
func (h *HttpResponseHeaders) ContentTypeText() *HttpResponseHeaders {
	return h.ContentType("text/plain")
}

func ContentTypeJson() HttpResponseHeaders {
	headers := HttpResponseHeaders{}
	headers.ContentTypeJson()
	return headers
}

func ContentTypeText() HttpResponseHeaders {
	headers := HttpResponseHeaders{}
	headers.ContentTypeText()
	return headers
}

// HttpResponse models the HttpTrigger out binding data.
// All fields are optional. The default value for Status is 200.
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
		Status  uint32              `json:"status"`
		Body    interface{}         `json:"body"`
		Headers HttpResponseHeaders `json:"headers"`
	}{
		h.Status, h.Body, h.Headers,
	})
}
