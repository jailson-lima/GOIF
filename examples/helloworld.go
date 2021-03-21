package main

import (
	"encoding/json"
	"github.com/JPereirax/GOIF/context"
	"github.com/JPereirax/GOIF/extensions"
	"github.com/JPereirax/GOIF/internal"
	"github.com/JPereirax/GOIF/route"
	"github.com/JPereirax/GOIF/types"
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

	directPosts()

	testPostRoute := v1Route.NewRoute("/post", http.MethodGet)
	testPostRoute.MultiStep(
		"log:Sending POST request to JSONPlaceholder",
		"direct:setbody",
		"http://jsonplaceholder.typicode.com/posts[method=post]",
		"direct:showposts",
	).Build()

	testGetRoute := v1Route.NewRoute("/showpost", http.MethodGet)
	testGetRoute.MultiStep(
		"log:Sending GET request to JSONPlaceholder",
		"http://jsonplaceholder.typicode.com/posts/1[method=get]",
		"direct:showposts",
	).Build()
}

func directPosts() {
	internal.From("direct:setbody").Processor(func(transport *types.HttpTransport) *types.HttpTransport {
		post, _ := json.Marshal(&Posts{
			UserId: 1,
			Title:  "GOIF",
			Body:   "Example GOIF POST request",
		})
		transport.Body = string(post)
		return transport
	}).End()

	internal.From("direct:showposts").Processor(func(transport *types.HttpTransport) *types.HttpTransport {
		extensions.SendJson(transport, &Posts{}, 200)
		return transport
	}).End()
}
