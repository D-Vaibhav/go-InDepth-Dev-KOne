package handlers

import (
	"fmt"
	"net/http"
)

func HomePageHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "this is homepage of our app")
}
