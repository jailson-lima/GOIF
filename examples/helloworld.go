package main

import (
	"GOIF/context"
	"GOIF/extensions"
	"GOIF/internal"
	"GOIF/route"
	"GOIF/types"
	"net/http"
)

type Posts struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

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

	directShowPosts()

	helloWorldRoute := v1Route.NewRoute("/helloworld", http.MethodGet)
	helloWorldRoute.Step("log:Starting hello world...")
	helloWorldRoute.MultiStep(
		"http://jsonplaceholder.typicode.com/posts/1[method=get]",
		"direct:showposts",
	)
	helloWorldRoute.Step("log:Finish hello world!")
	helloWorldRoute.Build()
}

func directShowPosts() {
	internal.From("direct:showposts").Processor(func(transport *types.HttpTransport) *types.HttpTransport {
		extensions.SendJson(transport, &Posts{}, 200)
		return transport
	}).End()
}
