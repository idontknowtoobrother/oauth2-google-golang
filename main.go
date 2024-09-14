package main

import (
	"log"
	"net/http"
	"os"

	"github.com/idontknowtoobrother/oauth2-google-golang/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("%v", err)
	}

	log.Println("Server closed!")
}
