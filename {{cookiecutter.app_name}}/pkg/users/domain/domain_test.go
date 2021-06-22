package domain_test

import (
	"testing"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		TestName string
		Name     string
		Surname  string
		Email    string
		Age      int
		Expected *entities.User
	}{
		{
			"Test create user",
			"Ruben",
			"Espinosa",
			"ruben@devaway.io",
			33,
			&entities.User{
				Name:     "Ruben",
				Surname:  "Espinosa",
				Username: "ruben.espinosa",
				Email:    "ruben@devaway.io",
				Age:      33,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			got := domain.NewUser(test.Name, test.Surname, test.Email, test.Age)
			if got.ID == "" {
				t.Errorf("Expected: %v, got: %v", test.Expected.ID, got.ID)
			}
			if got.Name != test.Expected.Name {
				t.Errorf("Expected: %v, got: %v", test.Expected.Name, got.Name)
			}
			if got.Surname != test.Expected.Surname {
				t.Errorf("Expected: %v, got: %v", test.Expected.Surname, got.Surname)
			}
			if got.Username != test.Expected.Username {
				t.Errorf("Expected: %v, got: %v", test.Expected.Username, got.Username)
			}
			if got.Email != test.Expected.Email {
				t.Errorf("Expected: %v, got: %v", test.Expected.Email, got.Email)
			}
			if got.Age != test.Expected.Age {
				t.Errorf("Expected: %v, got: %v", test.Expected.Age, got.Age)
			}
		})
	}
}
