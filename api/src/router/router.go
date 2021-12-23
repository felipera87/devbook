package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate creates a new router with all routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
