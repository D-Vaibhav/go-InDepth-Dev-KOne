package main

import (
	"net/http"
	"vaibhav/grpc/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.ListenAndServe(":8000", nil)
}
