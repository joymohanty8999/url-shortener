package handlers

import (
	"context"
	"net/http"
	"time"
	"url-shortener/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type RetrieveResponse struct {
	OriginalURL string `json:"original_url"`
}

func RetrieveURL(w http.ResponseWriter, r *http.Request) {

	//log.Println("RetrieveURL endpoint hit")

	vars := mux.Vars(r) // Extract the URL parameter from the request URL.
	shortURL := vars["shortURL"]

	// Query the database for the original URL corresponding to the short URL

	var entry models.URL
	filter := bson.M{"short_url": shortURL, "expiration": bson.M{"$gt": time.Now()}} // Check if the short URL exists and has not expired
	err := urlCollection.FindOne(context.TODO(), filter).Decode(&entry)
	if err != nil {
		http.Error(w, "URL not found or expired", http.StatusNotFound) // Return a 404 Not Found error if the short URL does not exist or has expired
		return
	}

	//log.Println("Redirecting:", shortURL, "➡", entry.OriginalURL)

	http.Redirect(w, r, entry.OriginalURL, http.StatusFound) // Redirect the user to the original URL
}
