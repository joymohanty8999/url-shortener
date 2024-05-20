package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"
)

var urlStore = make(map[string]models.URL)

const baseURL = "http://short.url/"
const expirationDuration = 2 * time.Minute //adding a time of 2 minutes per short url generated

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check if URL is already in use
	for shortURL, entry := range urlStore {
		if entry.OriginalURL == request.URL && time.Now().Before(entry.Expiration) {
			response := ShortenResponse{ShortURL: baseURL + shortURL}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	shortURL, err := utils.GenerateRandomBase62String(8)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expiration := time.Now().Add(expirationDuration)
	urlStore[shortURL] = models.URL{OriginalURL: request.URL, Expiration: expiration}

	response := ShortenResponse{ShortURL: baseURL + shortURL}
	json.NewEncoder(w).Encode(response)
}
