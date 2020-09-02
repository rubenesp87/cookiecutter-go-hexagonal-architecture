package entities

import "time"

// Post entity struct
type Post struct {
	ID      string    `json:"id"`
	Content string    `json:"content" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
	UserID  string    `json:"user_id" validate:"required"`
}
