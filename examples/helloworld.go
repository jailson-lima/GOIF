package main

import (
	"GOIF/context"
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

	helloWorldRoute := v1Route.NewRoute("/helloworld", http.MethodGet)
	helloWorldRoute.MultiStep("log:Testing 1", "log:Testing 2").Step("log:Testing 3")
	helloWorldRoute.Build()
}
