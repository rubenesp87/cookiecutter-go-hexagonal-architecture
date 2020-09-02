package inmemory

import (
	"fmt"

	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/posts/domain/entities"
)

// PostAdapter ...
type PostAdapter struct {
	Post *entities.Post
}

// NewInMemoryStorage adapter
func NewInMemoryStorage() *PostAdapter {
	return &PostAdapter{}
}

// Create inmemory post
func (pa *PostAdapter) Create(post *entities.Post) (*entities.Post, error) {
	pa.Post = post
	return post, nil
}

// Delete inmemory post
func (pa *PostAdapter) Delete(id string) error {
	pa.Post = nil
	return nil
}

// Get inmemory post
func (pa *PostAdapter) Get(id string) (*entities.Post, error) {
	if pa.Post == nil || pa.Post.ID != id {
		return &entities.Post{}, fmt.Errorf("Post not found")
	}
	return pa.Post, nil
}
