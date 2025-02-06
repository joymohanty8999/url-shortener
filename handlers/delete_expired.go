package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteExpiredURLs(w http.ResponseWriter, r *http.Request) {
	collection := utils.Client.Database("url-shortener").Collection("urls")

	// Ensure Go uses UTC time for comparison
	currentTime := time.Now().UTC()

	// Delete expired URLs based on expiration timestamp
	filter := bson.M{"expiration": bson.M{"$lte": currentTime}}
	result, err := collection.DeleteMany(r.Context(), filter)
	if err != nil {
		log.Printf("Error deleting expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted %v expired URLs", result.DeletedCount)))
}
