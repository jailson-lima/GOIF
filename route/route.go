package route

import (
	"log"
	"net/http"
)

type Rest struct {
	Id       string
	BasePath string
}

type Route struct {
	path   string
	method string
}

func (rest *Rest) NewRoute(path string, method string) *Route {
	return &Route{
		path:   rest.BasePath + path,
		method: method,
	}
}

func (route *Route) Build() {
	log.Printf("Registering route: %s", route.path)

	http.HandleFunc(route.path, func(response http.ResponseWriter, request *http.Request) {
		if request.Method != route.method {
			http.Error(response, "Method is not supported.", http.StatusMethodNotAllowed)
		}

		log.Printf("Received request from host '%s' to route '%s'.", request.RemoteAddr, request.Host + route.path)
	})
}