package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) { //ResponseWriter is a interface (not a star) and Request is a struct (so a star)
	g.l.Println("Goodbye World")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	//log.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "Goodbye %s", d)
}
