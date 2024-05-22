package main

import (
	"log"
	"net/http"
	"os"
	"url-shortener/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortURL}", handlers.RetrieveURL).Methods("GET")
	router.HandleFunc("/check", handlers.CheckURL).Methods("POST")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Starting server on port:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
