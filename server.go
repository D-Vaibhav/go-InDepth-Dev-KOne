package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"vaibhav/grpc/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	// creating oour own server with all property
	server := &http.Server{
		Addr:         ":8000",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// listening in go func, to avoid blocking
	go func() {
		// listening to the server not http
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// channel to recieve signal of type os.Signal
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt) //notifying if interrupted or killed
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	l.Println("Recieved termination signal with graceful shutdown", sig)

	// to gracefully shut-down
	shutdownContextTimeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(shutdownContextTimeout)
}
