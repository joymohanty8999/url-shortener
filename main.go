package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"url-shortener/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./front-end"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//serving index.html at root URL
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("front-end", "index.html"))
	}).Methods("GET")

	router.HandleFunc("/api/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
	}).Methods("GET")

	router.HandleFunc("/api/shorten", handlers.ShortenURL).Methods("POST")

	router.HandleFunc("/api/check", handlers.CheckURL).Methods("POST")

	router.HandleFunc("/api/urls", handlers.GetAllURLs).Methods("GET")

	router.HandleFunc("/api/delete-expired", handlers.DeleteExpiredURLs).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/{shortURL}", handlers.RetrieveURL).Methods("GET")

	// Handle preflight requests
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
