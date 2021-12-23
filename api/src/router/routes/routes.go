package routes

import "net/http"

// Route represents all routes on the API
type Route struct {
	URI          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}
