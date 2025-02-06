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

	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	var entry models.URL
	filter := bson.M{"short_url": shortURL, "expiration": bson.M{"$gt": time.Now()}}
	err := urlCollection.FindOne(context.TODO(), filter).Decode(&entry)
	if err != nil {
		http.Error(w, "URL not found or expired", http.StatusNotFound)
		return
	}

	//log.Println("Redirecting:", shortURL, "➡", entry.OriginalURL)

	http.Redirect(w, r, entry.OriginalURL, http.StatusFound)
}
