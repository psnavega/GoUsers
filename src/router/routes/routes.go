package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Struct all routes in API
type Route struct {
	URI               string
	Method            string
	Function          func(http.ResponseWriter, *http.Request)
	NeedAutentication bool
}

// Configure all routes inside users
func Configurate(r *mux.Router) *mux.Router {
	routes := routeUser

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
