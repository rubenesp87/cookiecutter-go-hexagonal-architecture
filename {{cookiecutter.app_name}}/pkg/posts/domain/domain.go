package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain/entities"
)

// Repository common functions to all repositories
type Repository interface {
	Create(post *entities.Post) (*entities.Post, error)
	Delete(id string) error
	Get(id string) (*entities.Post, error)
}

//NewPost returns curated user
func NewPost(content, userID string) *entities.Post {
	return &entities.Post{
		ID:      uuid.New().String(),
		Content: content,
		Date:    time.Now(),
		UserID:  userID,
	}
}
