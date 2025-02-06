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

	// Delete URLs where the expiration timestamp has passed
	result, err := collection.DeleteMany(r.Context(), bson.M{"expiration": bson.M{"$lt": time.Now()}})
	if err != nil {
		log.Printf("Error deleting expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted %v expired URLs", result.DeletedCount)))
}
