package application

import (
	"errors"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"
)

// Usecases represent the user business logic
type Usecases interface {
	RegisterUser(name, surname, email, password string, age int) (*entities.User, error)
	LoginUser(email, password string) (*entities.User, error)
	LogoutUser(id string) error
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

func (uu userUsecase) RegisterUser(name, surname, email, password string, age int) (*entities.User, error) {
	user := domain.NewUser(name, surname, email, password, age)
	return uu.userRepository.Create(user)
}

func (uu userUsecase) LoginUser(email, password string) (*entities.User, error) {
	user, err := uu.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if !domain.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Invalid password")
	}
	return user, nil
}

func (uu userUsecase) LogoutUser(id string) error {
	_, err := uu.userRepository.GetByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (uu userUsecase) DeleteUser(id string) error {
	return uu.userRepository.Delete(id)
}

func (uu userUsecase) GetUser(id string) (*entities.User, error) {
	return uu.userRepository.GetByID(id)
}
