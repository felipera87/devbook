package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, loginRoutes)
	routes = append(routes, publicationRoutes...)

	for _, route := range routes {
		if route.RequiresAuth {
			// middlewares works like nested functions
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Handler),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	return r
}
