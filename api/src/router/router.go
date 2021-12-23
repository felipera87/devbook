package router

import "github.com/gorilla/mux"

// Generate creates a new router with all routes configured
func Generate() *mux.Router {
	return mux.NewRouter()
}
