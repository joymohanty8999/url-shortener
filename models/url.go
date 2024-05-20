package models

import "time"

type URL struct {
	OriginalURL string
	Expiration  time.Time
}
