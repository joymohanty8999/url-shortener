package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var urlCollection *mongo.Collection

const baseURL = "https://snip-snip-go-2f69a42960b8.herokuapp.com/api/"
const expirationDuration = 15 * time.Second //adding a time of 2 hours per short url generated

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func init() {
	utils.InitDB()
	urlCollection = utils.GetCollection(utils.Client, "urls")
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {

	var request ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Remove any expired entries for the original URL
	filter := bson.M{"original_url": request.URL, "expiration": bson.M{"$lte": time.Now().UTC().Truncate(time.Millisecond)}}
	_, err := urlCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if a non-expired URL already exists
	var existingURL models.URL
	filter = bson.M{"original_url": request.URL, "expiration": bson.M{"$lte": time.Now().UTC().Truncate(time.Millisecond)}}
	err = urlCollection.FindOne(context.TODO(), filter).Decode(&existingURL)
	if err == nil {
		// Return the existing non-expired short URL
		response := ShortenResponse{ShortURL: baseURL + existingURL.ShortURL}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Generating new short url

	shortURL, err := utils.GenerateRandomBase62String(8)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expiration := time.Now().Add(expirationDuration).Truncate(time.Millisecond)
	newURL := models.URL{ShortURL: shortURL, OriginalURL: request.URL, Expiration: expiration}

	_, err = urlCollection.InsertOne(context.TODO(), newURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ShortenResponse{ShortURL: baseURL + shortURL}
	json.NewEncoder(w).Encode(response)
}
