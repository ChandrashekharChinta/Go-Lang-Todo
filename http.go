package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my Go server!")
}

func main() {

	// Step 1: Register routes
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/about", aboutHandler)

	// Step 2: Start server
	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

// The following happens:
// 1. Go starts executing main()

//    This is the entry point of your application.
// 2. Registers /hello

//    http.HandleFunc("/hello", helloHandler) stores the route mapping.
// 3. Registers /about

//    http.HandleFunc("/about", aboutHandler) stores another route mapping.
// 4. Starts the server

//    http.ListenAndServe(":8080", nil) starts listening.
// 5. Browser sends a request

//    You visit http://localhost:8080/hello.
// 6. Go selects the handler

//    It finds /hello and calls helloHandler.
// 7. Handler writes the response

//    fmt.Fprint(w, "Hello, World!").
// 8. Browser receives the response

//    It displays Hello, World!.
