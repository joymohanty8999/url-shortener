package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteExpiredURLs(w http.ResponseWriter, r *http.Request) {
	collection := utils.Client.Database("url_shortener").Collection("urls")

	// Get current UTC time and round to match MongoDB precision
	currentTime := time.Now().UTC().Truncate(time.Millisecond)

	// Query MongoDB for expired URLs and store them in a slice or array
	var expiredURLs []models.URL
	cursor, err := collection.Find(r.Context(), bson.M{"expiration": bson.M{"$lte": currentTime}})
	if err != nil {
		log.Printf("Error finding expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError) // Return a 500 Internal Server Error if the query fails
		return
	}
	defer cursor.Close(r.Context())

	if err = cursor.All(r.Context(), &expiredURLs); err != nil {
		log.Printf("Error decoding expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Log the expired URLs
	if len(expiredURLs) == 0 {
		log.Println("No expired URLs found")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No expired URLs found"))
		return
	} else {
		log.Println("Expired URLs found, deleting now:", expiredURLs)
	}

	// Attempt to delete expired URLs
	result, err := collection.DeleteMany(r.Context(), bson.M{"expiration": bson.M{"$lte": currentTime}})
	if err != nil {
		log.Printf("Error deleting expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Log the number of deleted URLs and return a success response

	log.Printf("Deleted %v expired URLs", result.DeletedCount)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted %v expired URLs", result.DeletedCount)))
}
