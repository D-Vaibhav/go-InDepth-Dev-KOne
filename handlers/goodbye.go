package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Requested for goodbye page")
	requestBodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Unable to load goodbye page", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Requested for goodbye page with recieved body is: %s", requestBodyData)
}
