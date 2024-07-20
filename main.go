package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"url-shortener/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./front-end"))
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))

	//log.Println("Setting up routes")

	//serving index.html at root URL
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("front-end", "index.html"))
		//log.Println("Root endpoint hit")
	}).Methods("GET")

	router.HandleFunc("/api/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
		//log.Println("Favicon endpoint hit")
	}).Methods("GET")

	router.HandleFunc("/api/shorten", handlers.ShortenURL).Methods("POST")
	//log.Println("Registered /shorten endpoint")

	router.HandleFunc("/api/check", handlers.CheckURL).Methods("POST")
	//log.Println("Registered /check endpoint")

	router.HandleFunc("/api/urls", handlers.GetAllURLs).Methods("GET")
	//log.Println("Registered /urls endpoint")

	router.HandleFunc("/api/delete-expired", handlers.DeleteExpiredURLs).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/{shortURL}", handlers.RetrieveURL).Methods("GET")
	//log.Println("Registered /{shortURL} endpoint")

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
