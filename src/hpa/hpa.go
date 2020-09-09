package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
)

var x = 0.0001

func Looping() (r string) {
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}

	return "Code.education Rocks!"
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	msg := Looping()
	fmt.Fprintf(w, msg)
}

func main() {
	// register index function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	// use PORT environment variable, or default to 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
