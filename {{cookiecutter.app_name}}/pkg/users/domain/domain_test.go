package domain_test

import (
	"testing"

	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/domain"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/users/domain/entities"
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
			if *got != *test.Expected {
				t.Errorf("Expected: %v, got: %v", test.Expected, got)
			}
		})
	}
}
