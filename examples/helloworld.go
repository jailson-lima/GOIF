package main

import (
	"GOIF/context"
	"GOIF/internal"
	"GOIF/route"
	"GOIF/types"
	"net/http"
)

func main() {
	properties := types.Properties{
		Host:           "127.0.0.1",
		Port:           8080,
		RegisterRoutes: registerRoutes,
	}

	context.Start(properties)
}

func registerRoutes() {
	v1Route := &route.Rest{
		Id:       "V1 API",
		BasePath: "/v1/api",
	}

	setupShowNameFunction()

	helloWorldRoute := v1Route.NewRoute("/helloworld", http.MethodGet)
	helloWorldRoute.MultiStep("log:Starting hello world...", "direct:showname").Step("log:Finish hello world!")
	helloWorldRoute.Build()
}

func setupShowNameFunction() {
	internal.From("direct:showname").Processor(func(transport types.HttpTransport) types.HttpTransport {
		transport.Response.Write([]byte("Hello World!"))
		return transport
	}).End()
}
