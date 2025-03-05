package models

import "time"

type Article struct {
	Slug        string
	Title       string
	Description string
	IsDraft     bool
	Categories  []string
	PublishedAt time.Time
	ModifiedAt  time.Time
	Content     []byte
}
