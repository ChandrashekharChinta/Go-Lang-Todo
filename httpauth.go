package main

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("X-Auth-Token")
		if token != validToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

	// w http.ResponseWriter: Used to send a response back to the client.
	// r *http.Request: Contains information about the incoming request.
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are authorized!")

}

const validToken = "secret"

// HandleFunc()			Registers a handler function directly
// Handle()				Registers an http.Handler
// NewServeMux()		Creates the router
// AuthMiddleware()		Checks authentication
// ServeHTTP()			Executes the next handler

func SetupServer() http.Handler {

	mux := http.NewServeMux()

	//public endpoin
	mux.HandleFunc("/hello", HelloHandler)

	secureRoute := http.HandlerFunc(SecureHandler)
	mux.Handle("/secure", AuthMiddleware(secureRoute))

	return mux
}

func main() {

	server := SetupServer()

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", server)

	if err != nil {
		fmt.Println("Server error:", err)
	}

}

// http.Handler					Interface for handling HTTP requests
// http.HandlerFunc				Converts a function into an HTTP handler
// http.ResponseWriter			Writes the HTTP response
// http.Request				    Contains incoming request information
// http.NewServeMux()			Creates a router
// http.StatusUnauthorized		HTTP status code 401
