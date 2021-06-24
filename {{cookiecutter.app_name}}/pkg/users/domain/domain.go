package domain

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"
	"golang.org/x/crypto/bcrypt"
)

// Repository common functions to all repositories
type Repository interface {
	Create(user *entities.User) (*entities.User, error)
	Delete(id string) error
	GetByID(id string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
}

//HashPassword encrypt user password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash check user password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//NewUser returns curated user
func NewUser(name, surname, email, password string, age int) *entities.User {
	username := fmt.Sprintf("%s.%s", strings.ToLower(name), strings.ToLower(surname))
	hash, _ := HashPassword(password)

	return &entities.User{
		ID:       uuid.New().String(),
		Name:     name,
		Surname:  surname,
		Username: username,
		Email:    email,
		Password: hash,
		Age:      age,
	}
}
