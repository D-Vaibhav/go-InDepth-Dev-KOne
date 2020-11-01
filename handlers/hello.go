package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Requested for home page")
	requestBodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops, its an bad Request", http.StatusBadRequest)
		return
	}
	// log.Fatalf("Failed to load: %v", err)
	fmt.Fprintf(rw, "Requested for home page with passed data: %s", requestBodyData)
}
