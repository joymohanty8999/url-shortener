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
	collection := utils.Client.Database("url-shortener").Collection("urls")

	// Get current UTC time
	currentTime := time.Now().UTC()

	// Find all expired URLs (for debugging)
	var expiredURLs []models.URL
	cursor, err := collection.Find(r.Context(), bson.M{"expiration": bson.M{"$lte": currentTime}})
	if err != nil {
		log.Printf("Error finding expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
	} else {
		log.Println("Expired URLs:", expiredURLs)
	}

	// Attempt to delete expired URLs
	result, err := collection.DeleteMany(r.Context(), bson.M{"expiration": bson.M{"$lte": currentTime}})
	if err != nil {
		log.Printf("Error deleting expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Deleted %v expired URLs", result.DeletedCount)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted %v expired URLs", result.DeletedCount)))
}
