package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	// HandleFunc is a convenience method to register a function to a path on DefaultServeMux
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		//log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")

	})

	http.ListenAndServe("192.168.122.17:9090", nil)
}
