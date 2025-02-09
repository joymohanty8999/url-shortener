package main

import (
	"fmt"
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

	// Load environment variables from .env file -> This allows us to securely store sensitive information such as our MongoDB URI connection string
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables") // falls back to system environment variables if .env file is not found to ensure app is always able to run
	}

	// Print environment variable for debugging
	mongoURI := os.Getenv("MONGODB_URI")
	fmt.Println("MONGODB_URI:", mongoURI) // Debugging line

	if mongoURI == "" {
		log.Fatal("Error: MONGODB_URI environment variable not set")
	}

	// Initializing Gorilla Mux router

	router := mux.NewRouter()

	// serving front-end files

	fs := http.FileServer(http.Dir("./front-end"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//serving index.html at root URL
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("front-end", "index.html"))
	}).Methods("GET")

	//serving urls.html at /urls.html

	router.HandleFunc("/urls.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("front-end", "urls.html"))
	}).Methods("GET")

	router.HandleFunc("/api/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
	}).Methods("GET")

	//API Endpoints

	router.HandleFunc("/api/shorten", handlers.ShortenURL).Methods("POST") // shortens the user-provided URL to a base62 encoded link and stores it in MongoDB

	router.HandleFunc("/api/check", handlers.CheckURL).Methods("POST") // checks if the user-provided URL is already in the database

	router.HandleFunc("/api/urls", handlers.GetAllURLs).Methods("GET") // retrieves all shortened URLs stored in the database

	router.HandleFunc("/api/delete-expired", handlers.DeleteExpiredURLs).Methods("DELETE", "OPTIONS") // automatically deletes expired URLs from the database

	router.HandleFunc("/api/{shortURL}", handlers.RetrieveURL).Methods("GET") // When a user visits a shortened URL, this endpoint retrieves the original URL from the database and redirects the user to the original URL

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
