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

	cursor, err := utils.Client.Database("url_shortener").Collection("urls").Find(ctx, bson.M{})

	if err != nil {
		log.Printf("Error finding URLs: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)

	var urls []models.URL

	if err = cursor.All(ctx, &urls); err != nil {
		log.Printf("Error decoding URLs: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(urls) == 0 {
		http.Error(w, "No URLs found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)

}
