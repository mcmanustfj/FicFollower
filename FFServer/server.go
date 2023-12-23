// A basic web server in Go

package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func resetdb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Resetting DB")

	err := Recreate_db()
	if err != nil {
		fmt.Fprintf(w, "\nError: %s", err)
		return
	}

	fmt.Fprintf(w, "\nSuccess")
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/resetdb", resetdb)

	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	http.HandleFunc("/", home)

	// ListenAndServe listens on the TCP network address addr and
	// then calls Serve with handler to handle requests on incoming connections.

	// fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
