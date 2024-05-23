package models

import "time"

type URL struct {
	ShortURL    string    `json:"short_url" bson:"short_url"`
	OriginalURL string    `json:"original_url" bson:"original_url"`
	Expiration  time.Time `json:"expiration" bson:"expiration"`
}
