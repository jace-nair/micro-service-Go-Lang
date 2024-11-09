package main

import (
	"context"
	"log"
	"micro-service-app/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	// Manually configuring the server
	s := &http.Server{
		Addr:         "192.168.122.17:9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Handl Listen and Serve in go func to avoid blocking the service
	go func() {
		// Listen and serve
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigChan := make(chan os.Signal, 1)
	// Broadcast a message on this channel wheneven an operating system (os) Interrupt or Kill command is received
	signal.Notify(sigChan, os.Interrupt)
	// Register for signal
	signal.Notify(sigChan, os.Kill)

	// Block until any signal is received
	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	// Once the message is consumed, create a context for graceful shutdown. tc is time context
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Graceful server shutdown with tc
	s.Shutdown(tc)

}
