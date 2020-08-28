package domain

import (
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/posts/domain/entities"
)

// Repository common functions to all repositories
type Repository interface {
	Create(user *entities.Post) error
	Delete(id string) (int64, error)
	Get(id string) (*entities.Post, error)
}

//NewPost returns curated user
func NewPost(content, date string) *entities.Post {

	return &entities.Post{
		Content: content,
		Date:    date,
		UserID:  "",
	}
}
