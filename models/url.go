package models

import "time"

type URL struct {
	ShortURL    string    `bson:"short_url"`
	OriginalURL string    `bson:"original_url"`
	Expiration  time.Time `bson:"expiration"`
}
