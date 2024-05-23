package handlers

import (
	"context"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	shortURL := params["shortURL"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var url models.URL
	collection := utils.Client.Database("url_shortener").Collection("urls")
	err := collection.FindOne(ctx, bson.M{"short_url": shortURL}).Decode(&url)

	if err != nil {
		http.Error(w, "URL not found or expired", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)

}
