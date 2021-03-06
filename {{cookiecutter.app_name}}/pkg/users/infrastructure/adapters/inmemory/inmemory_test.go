package inmemory_test

import (
	"testing"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/infrastructure/adapters/inmemory"
)

type mockUserAdapter struct {
	Error error
	User  *entities.User
}

func TestCreate(t *testing.T) {
	tests := []struct {
		TestName      string
		User          *entities.User
		Expected      *entities.User
		ExpectedError error
	}{
		{
			"Test create user",
			&entities.User{
				ID:       "1234",
				Name:     "Ruben",
				Surname:  "Espinosa",
				Username: "ruben.espinosa",
				Email:    "ruben@devaway.io",
				Age:      33,
			},
			&entities.User{
				ID:       "1234",
				Name:     "Ruben",
				Surname:  "Espinosa",
				Username: "ruben.espinosa",
				Email:    "ruben@devaway.io",
				Age:      33,
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			userAdapter := inmemory.UserAdapter{}
			user, err := userAdapter.Create(test.User)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *user != *test.Expected {
				t.Errorf("Expected: %v, got: %v", test.Expected, user)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		User          *entities.User
		Expected      *entities.User
		ExpectedError error
	}{
		{
			"Test get user",
			"1234",
			&entities.User{
				ID:       "1234",
				Name:     "Ruben",
				Surname:  "Espinosa",
				Username: "ruben.espinosa",
				Email:    "ruben@devaway.io",
				Age:      33,
			},
			&entities.User{
				ID:       "1234",
				Name:     "Ruben",
				Surname:  "Espinosa",
				Username: "ruben.espinosa",
				Email:    "ruben@devaway.io",
				Age:      33,
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			userAdapter := inmemory.UserAdapter{
				User: test.User,
			}
			user, err := userAdapter.GetByID(test.ID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *user != *test.Expected {
				t.Errorf("Expected: %v, got: %v", test.Expected, user)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		User          *entities.User
		ExpectedError error
	}{
		{
			"Test delete user",
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
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			userAdapter := inmemory.UserAdapter{
				User: test.User,
			}
			err := userAdapter.Delete("1234")
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
		})
	}
}
