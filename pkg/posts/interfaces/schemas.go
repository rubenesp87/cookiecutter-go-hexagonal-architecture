package interfaces

// CreatePostRequest echo struct
type CreatePostRequest struct {
	Content string `json:"content" validate:"required"`
}
