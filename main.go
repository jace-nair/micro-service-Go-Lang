package main

import (
	"log"
	"micro-service-app/handlers"
	"net/http"
	"os"
)

func main() {

	// HandleFunc is a convenience method to register a function to a path on DefaultServeMux
	// Instead of DefaltServeMux, create a custom servemux server

	// Logger reference
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Create handler references
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// Create a new servemux server
	sm := http.NewServeMux()

	// Register the new handlers with servemux server
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// Bind address to the newly creater servemux server
	http.ListenAndServe("192.168.122.17:9090", sm)
}
