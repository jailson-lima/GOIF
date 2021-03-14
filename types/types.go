package types

import (
	"net/http"
)

type Properties struct {
	Host           string
	Port           int
	RegisterRoutes func()
}

type HttpTransport struct {
	Request *http.Request
	Response http.ResponseWriter
}

type InternalFunction struct {
	Id         string
	Processors []func(transport HttpTransport) HttpTransport
}