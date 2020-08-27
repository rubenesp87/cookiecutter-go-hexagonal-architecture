package domain

import (
	"fmt"
	"strings"

	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/domain/entities"
)

// Repository common functions to all repositories
type Repository interface {
	Create(user *entities.User) error
	Delete(id string) (int64, error)
	Get(id string) (*entities.User, error)
}

//NewUser returns curated user
func NewUser(name, surname, email string, age int) *entities.User {
	username := fmt.Sprintf("%s.%s", strings.ToLower(name), strings.ToLower(surname))

	return &entities.User{
		Name:     name,
		Surname:  surname,
		Username: username,
		Email:    email,
		Age:      age,
	}
}
