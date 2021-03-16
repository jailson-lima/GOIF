package types

import "net/http"

type (
	Properties struct {
		Host           string
		Port           int
		RegisterRoutes func()
	}

	HttpTransport struct {
		Request *http.Request
		Response http.ResponseWriter

		Body     string
	}

	ContextRequest struct {
		Component Component

		HttpTransport *HttpTransport
	}

	Component struct {
		Schema    string
		Uri       string

		Function func(ctx ContextRequest) (*HttpTransport, error)
	}

	InternalFunction struct {
		Id         string
		Processors []func(transport *HttpTransport) *HttpTransport
	}

	ErrorResponse struct {
		Status  int
		Message string
	}
)