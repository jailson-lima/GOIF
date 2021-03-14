package route

import (
	"net/http"
	"testing"
)

var rest = &Rest{
	Id:       "ping",
	BasePath: "/ping",
}

var pongRoute = rest.NewRoute("/pong", http.MethodGet)

func TestRest_NewRoute(t *testing.T) {
	if pongRoute == nil {
		t.Error("Rest#NewRoute() failed, expected return Route{}.")
	} else {
		t.Logf("Rest#NewRoute() success, expected Route{}, got %v.", pongRoute)
	}
}