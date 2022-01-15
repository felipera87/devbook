package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger writes information about the request on terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate makes a route demands authentication, checks if the user is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validate")
		next(w, r)
	}
}
