package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all routes on the API
type Route struct {
	URI          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// Configure receives a not configured router and returns it with all routes
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
