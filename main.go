package main

import (
	"log"
	"net/http"
	"url-shortener/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortURL}", handlers.RetrieveURL).Methods("GET")
	router.HandleFunc("/check", handlers.CheckURL).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
