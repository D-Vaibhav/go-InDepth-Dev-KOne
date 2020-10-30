package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HomePageHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Requested for getBody page")
	fmt.Fprintf(rw, "this is homepage of our app")
}

func GetBodyHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Requested for getBody page")

	requestBodyData, _ := ioutil.ReadAll(r.Body)
	// log.Fatalf("Failed to load: %v", err)
	fmt.Fprintf(rw, "this is getBody page with passed data: %s", requestBodyData)
}
