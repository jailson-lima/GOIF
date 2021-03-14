package route

import (
	"GOIF/component"
	"GOIF/types"
	"log"
	"net/http"
	"strings"
)

type Rest struct {
	Id       string
	BasePath string
}

type Route struct {
	path   string
	method string
	steps  []component.Component
}

func (rest *Rest) NewRoute(path string, method string) *Route {
	return &Route{
		path:   rest.BasePath + path,
		method: method,
	}
}

func (route *Route) MultiStep(steps ...string) *Route {
	for _, step := range steps {
		route.Step(step)
	}
	return route
}

func (route *Route) Step(step string) *Route {
	compound := strings.Split(step, ":")

	for _, components := range component.RegisteredComponents {
		if components.Schema == compound[0] {
			components.Uri = compound[1]
			route.steps = append(route.steps, *components)
		}
	}

	return route
}

func (route *Route) Build() {
	log.Printf("Registering route: %s", route.path)

	http.HandleFunc(route.path, func(response http.ResponseWriter, request *http.Request) {
		if request.Method != route.method {
			http.Error(response, "Method is not supported.", http.StatusMethodNotAllowed)
		}

		log.Printf("Received request from host '%s' to route '%s'.", request.RemoteAddr, request.Host + route.path)

		for _, step := range route.steps {
			context := component.ContextRequest{
				Component: step,
				HttpTransport: types.HttpTransport{
					Request:  request,
					Response: response,
				},
			}

			transportProperty := step.Function(context)
			request, response = transportProperty.Request, transportProperty.Response
			if response == nil {
				break
			}
		}
	})
}