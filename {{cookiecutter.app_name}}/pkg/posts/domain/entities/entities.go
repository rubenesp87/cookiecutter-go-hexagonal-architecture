package entities

// Post entity struct
type Post struct {
	Content string `json:"content" validate:"required"`
	Date    string `json:"date" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
}
