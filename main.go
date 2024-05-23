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

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the URL Shortener API. Use /shorten, /{shortURL}, /check, /urls endpoints"))
		log.Println("Root Endpoint hit")
	}).Methods("GET")

	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
		log.Println("Favicon endpoint hit")
	}).Methods("GET")

	router.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortURL}", handlers.RetrieveURL).Methods("GET")
	router.HandleFunc("/check", handlers.CheckURL).Methods("POST")

	router.HandleFunc("/urls", handlers.GetAllURLs).Methods("GET")
	log.Println("Registered /urls endpoint")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Router configured successfully. Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
