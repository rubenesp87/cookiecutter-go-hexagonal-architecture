package inmemory

import (
	"fmt"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"
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
func (ua *UserAdapter) Create(user *entities.User) (*entities.User, error) {
	ua.User = user
	return user, nil
}

// Delete inmemory user
func (ua *UserAdapter) Delete(id string) error {
	ua.User = nil
	return nil
}

// Get inmemory user
func (ua *UserAdapter) Get(id string) (*entities.User, error) {
	if ua.User == nil || ua.User.ID != id {
		return &entities.User{}, fmt.Errorf("User not found")
	}
	return ua.User, nil
}
