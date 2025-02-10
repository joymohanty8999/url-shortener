package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"url-shortener/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// The request struct for the check endpoint which holds the URL to check

type CheckRequest struct {
	URL string `json:"url"`
}

// The repsonse struct for the check endpoint

type CheckResponse struct {
	Exists   bool   `json:"exists"`
	ShortURL string `json:"short_url,omitempty"`
	Expired  *bool  `json:"expired,omitempty"`
}

func CheckURL(w http.ResponseWriter, r *http.Request) {

	var request CheckRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingURL models.URL
	filter := bson.M{"original_url": request.URL}
	opts := options.FindOne().SetSort(bson.D{{Key: "expiration", Value: -1}}) // Gets the most recent entry for the URL
	err := urlCollection.FindOne(context.TODO(), filter, opts).Decode(&existingURL)
	if err != nil {
		response := CheckResponse{Exists: false}
		json.NewEncoder(w).Encode(response)
		return
	}

	expired := time.Now().After(existingURL.Expiration)
	response := CheckResponse{
		Exists:   true,
		ShortURL: baseURL + existingURL.ShortURL,
		Expired:  &expired,
	}
	json.NewEncoder(w).Encode(response)
}
