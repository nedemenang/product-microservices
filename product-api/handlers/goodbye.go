package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}


func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func(g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter is an interface, that is why where aren't using a pointer
	// http.Request is a struct that why where using a pointer
	rw.Write([]byte("Bye"))
}