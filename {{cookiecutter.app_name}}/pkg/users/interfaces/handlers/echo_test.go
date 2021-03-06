package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/users/domain/entities"

	"github.com/labstack/echo/v4"
)

type MockUseCases struct {
	Error error
}

func (m MockUseCases) RegisterUser(name, surname, email, password string, age int) (*entities.User, error) {
	return &entities.User{}, m.Error
}

func (m MockUseCases) LoginUser(email, password string) (*entities.User, error) {
	return &entities.User{}, m.Error
}

func (m MockUseCases) LogoutUser(id string) error {
	return m.Error
}

func (m MockUseCases) GetUser(id string) (*entities.User, error) {
	return &entities.User{}, m.Error
}

func (m MockUseCases) DeleteUser(id string) error {
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
		Error         error
		ExpectedError error
	}{
		{
			TestName:      "Test Create user OK",
			Name:          "Ruben",
			Surname:       "Espinosa",
			Email:         "ruben@devaway.io",
			Password:      "123456789",
			Age:           33,
			Error:         nil,
			ExpectedError: nil,
		},
		{
			TestName:      "Test Create user KO",
			Name:          "Anthony",
			Surname:       "Smith",
			Email:         "ant@sm.com",
			Password:      "123456789",
			Age:           25,
			Error:         fmt.Errorf("error"),
			ExpectedError: ErrCreatingUser,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			body := fmt.Sprintf(`{"name": "%s", "surname": "%s", "email": "%s", "password": "%s", "age": %v}`,
				test.Name,
				test.Surname,
				test.Email,
				test.Password,
				test.Age,
			)
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			fakeHandler := APIHandler{
				Usecases: MockUseCases{
					Error: test.Error,
				},
			}

			err := fakeHandler.RegisterUser(c)
			if err != test.ExpectedError {
				t.Errorf("got %v, expected %v", err, test.ExpectedError)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		Error         error
		ExpectedError error
	}{
		{
			TestName:      "Test Get user OK",
			ID:            "1234",
			Error:         nil,
			ExpectedError: nil,
		},
		{
			TestName:      "Test Get user KO",
			ID:            "1234",
			Error:         fmt.Errorf("error"),
			ExpectedError: ErrGettingUser,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/users/%v", test.ID),
				strings.NewReader(""),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			fakeHandler := &APIHandler{
				Usecases: MockUseCases{
					Error: test.Error,
				},
			}

			err := fakeHandler.GetUser(c)
			if err != test.ExpectedError {
				t.Errorf("got %v, expected %v", err, test.ExpectedError)
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
			TestName:      "Test Delete user OK",
			ID:            "existing_id",
			Error:         nil,
			ExpectedError: nil,
		},
		{
			TestName:      "Test Delete user KO",
			ID:            "non_existing_id",
			Error:         fmt.Errorf("error"),
			ExpectedError: ErrDeletingUser,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%v", test.ID), strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			fakeHandler := APIHandler{
				Usecases: MockUseCases{
					Error: test.Error,
				},
			}
			if err := fakeHandler.DeleteUser(c); err != test.ExpectedError {
				t.Errorf("got %v, expected %v", err, test.ExpectedError)
			}
		})
	}
}
