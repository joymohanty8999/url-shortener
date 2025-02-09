package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllURLs(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Prevents the database query from running indefinitely.
	defer cancel()

	collection := utils.Client.Database("url_shortener").Collection("urls") // Attempting to access the "urls" collection in the "url_shortener" database.
	if collection == nil {
		http.Error(w, "Database collection not found", http.StatusInternalServerError)
		return
	}

	cursor, err := collection.Find(ctx, bson.M{}) // Retrieving all documents in the "urls" collection.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var urls []models.URL

	// we iterate through the cursor and decode each document into a URL struct

	if err = cursor.All(ctx, &urls); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If no URLs are found, we return a 404 Not Found error instead of an empty array.

	if len(urls) == 0 {
		http.Error(w, "No URLs found", http.StatusNotFound)
		return
	}

	// Set the Content-Type header to application/json and encode the URLs array into JSON format before sending it as the response.

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}
