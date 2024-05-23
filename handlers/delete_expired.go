package handlers

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteExpiredURLs(w http.ResponseWriter, r *http.Request) {
	collection := utils.Client.Database("url-shortener").Collection("urls")

	result, err := collection.DeleteMany(r.Context(), bson.M{"expired": true})
	if err != nil {
		log.Printf("Error deleting expired URLs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted %v expired URLs", result.DeletedCount)))
}
