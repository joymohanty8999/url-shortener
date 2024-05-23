package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllURLs(w http.ResponseWriter, r *http.Request) {

	log.Println("GetAllURLs endpoint hit")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Attempting to find URLs in the database")
	collection := utils.Client.Database("url_shortener").Collection("urls")
	if collection == nil {
		log.Println("Collection is nil")
		http.Error(w, "Database collection not found", http.StatusInternalServerError)
		return
	}

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error finding URLs: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var urls []models.URL
	log.Println("Attempting to decode URLs")
	if err = cursor.All(ctx, &urls); err != nil {
		log.Printf("Error decoding URLs: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(urls) == 0 {
		log.Println("No URLs found in the database")
		http.Error(w, "No URLs found", http.StatusNotFound)
		return
	}

	log.Printf("Returning %d URLs\n", len(urls))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}
