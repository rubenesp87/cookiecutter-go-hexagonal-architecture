package inmemory

import (
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/domain/entities"
)

// UserAdapter ...
type UserAdapter struct {
	User *entities.User
}

// NewInMemoryStorage adapter
func NewInMemoryStorage() *UserAdapter {
	return &UserAdapter{}
}

// Create inmemory user
func (ua *UserAdapter) Create(user *entities.User) error {
	ua.User = user
	return nil
}

// Delete inmemory user
func (ua *UserAdapter) Delete(id string) error {
	ua.User = nil
	return nil
}

// Get inmemory user
func (ua *UserAdapter) Get(id string) (*entities.User, error) {
	return ua.User, nil
}
