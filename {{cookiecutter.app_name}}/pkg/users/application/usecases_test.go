package application

import (
	"errors"
	"testing"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"
)

var errCreatingUser = errors.New("Error creating new User")
var errGetingUser = errors.New("Error geting a User")
var errDeletingUser = errors.New("Error deleting a User")

type mockUserRepository struct {
	User  *entities.User
	Error error
}

func (m mockUserRepository) Create(user *entities.User) (*entities.User, error) {
	return m.User, m.Error
}

func (m mockUserRepository) GetByID(id string) (*entities.User, error) {
	return m.User, m.Error
}

func (m mockUserRepository) GetByEmail(email string) (*entities.User, error) {
	return m.User, m.Error
}

func (m mockUserRepository) Delete(id string) error {
	return m.Error
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		TestName      string
		Name          string
		Surname       string
		Email         string
		Password      string
		Age           int
		User          *entities.User
		Error         error
		ExpectedError error
	}{
		{
			"Test create user OK",
			"Ruben",
			"Espinosa",
			"ruben@devaway.io",
			"123456789",
			33,
			&entities.User{
				ID:       "1234",
				Name:     "Ruben",
				Surname:  "Espinosa",
				Username: "ruben.espinosa",
				Email:    "ruben@devaway.io",
				Password: "123456789",
				Age:      33,
			},
			nil,
			nil,
		},
		{
			"Test create user KO",
			"Ruben",
			"Espinosa",
			"ruben@devaway.io",
			"123456789",
			33,
			&entities.User{},
			errCreatingUser,
			errCreatingUser,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {

			fakeUseCases := userUsecase{
				userRepository: mockUserRepository{
					User:  test.User,
					Error: test.Error,
				},
			}
			user, err := fakeUseCases.RegisterUser(test.Name, test.Surname, test.Email, test.Password, test.Age)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *user != *test.User {
				t.Errorf("Expected: %v, got: %v", test.User, user)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		User          *entities.User
		Error         error
		ExpectedError error
	}{
		{
			"Test get user OK",
			"1234",
			&entities.User{
				ID:       "1234",
				Name:     "Ruben",
				Surname:  "Espinosa",
				Username: "ruben.espinosa",
				Email:    "ruben@devaway.io",
				Age:      33,
			},
			nil,
			nil,
		},
		{
			"Test get user KO",
			"1234",
			&entities.User{},
			errGetingUser,
			errGetingUser,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {

			fakeUseCases := userUsecase{
				userRepository: mockUserRepository{
					User:  test.User,
					Error: test.Error,
				},
			}
			user, err := fakeUseCases.GetUser(test.ID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *user != *test.User {
				t.Errorf("Expected: %v, got: %v", test.User, user)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		Error         error
		ExpectedError error
	}{
		{
			"Test delete user OK",
			"1234",
			nil,
			nil,
		},
		{
			"Test delete user KO",
			"1234",
			errDeletingUser,
			errDeletingUser,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {

			fakeUseCases := userUsecase{
				userRepository: mockUserRepository{
					Error: test.Error,
				},
			}
			err := fakeUseCases.DeleteUser(test.ID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
		})
	}
}
