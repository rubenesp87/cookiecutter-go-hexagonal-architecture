package application

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"
)

// Usecases represent the user business logic
type Usecases interface {
	CreateUser(name, surname, email string, age int) (*entities.User, error)
	DeleteUser(id string) error
	GetUser(id string) (*entities.User, error)
}

type userUsecase struct {
	userRepository domain.Repository
}

// NewUserUsecase will create new an userUsecase of userRepository
func NewUserUsecase(ur domain.Repository) Usecases {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu userUsecase) CreateUser(name, surname, email string, age int) (*entities.User, error) {
	user := domain.NewUser(name, surname, email, age)
	return uu.userRepository.Create(user)
}

func (uu userUsecase) DeleteUser(id string) error {
	return uu.userRepository.Delete(id)
}

func (uu userUsecase) GetUser(id string) (*entities.User, error) {
	return uu.userRepository.Get(id)
}