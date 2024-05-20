package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type RetrieveResponse struct {
	OriginalURL string `json:"original_url"`
}

func RetrieveURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	entry, ok := urlStore[shortURL]
	if !ok || time.Now().After(entry.Expiration) {
		http.Error(w, "URL not found or expired", http.StatusNotFound)
		return
	}

	response := RetrieveResponse{OriginalURL: entry.OriginalURL}
	json.NewEncoder(w).Encode(response)
}
