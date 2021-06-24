package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain/entities"

	"github.com/labstack/echo/v4"
)

type MockUseCases struct {
	Error error
}

func (m MockUseCases) CreatePost(content, userID string) (*entities.Post, error) {
	return &entities.Post{}, m.Error
}

func (m MockUseCases) GetPost(id string) (*entities.Post, error) {
	return &entities.Post{}, m.Error
}

func (m MockUseCases) DeletePost(id string) error {
	return m.Error
}

func TestCreatePost(t *testing.T) {

	tests := []struct {
		TestName      string
		Content       string
		Error         error
		ExpectedError error
	}{
		{
			TestName:      "Test Create post OK",
			Content:       "My first post",
			Error:         nil,
			ExpectedError: nil,
		},
		{
			TestName:      "Test Create post KO",
			Content:       "My firts failed post",
			Error:         fmt.Errorf("error"),
			ExpectedError: ErrCreatingPost,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			body := fmt.Sprintf(`{"content": "%s"}`,
				test.Content,
			)
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			fakeHandler := APIHandler{
				Usecases: MockUseCases{
					Error: test.Error,
				},
			}

			err := fakeHandler.CreatePost(c)
			if err != test.ExpectedError {
				t.Errorf("got %v, expected %v", err, test.ExpectedError)
			}
		})
	}
}

func TestGetPost(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		UserID        string
		Error         error
		ExpectedError error
	}{
		{
			TestName:      "Test Get post OK",
			ID:            "1234",
			UserID:        "5678",
			Error:         nil,
			ExpectedError: nil,
		},
		{
			TestName:      "Test Get post KO",
			ID:            "1234",
			UserID:        "5678",
			Error:         fmt.Errorf("error"),
			ExpectedError: ErrGettingPost,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/posts/%v/%v", test.ID, test.UserID),
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

			err := fakeHandler.GetPost(c)
			if err != test.ExpectedError {
				t.Errorf("got %v, expected %v", err, test.ExpectedError)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		UserID        string
		Error         error
		ExpectedError error
	}{
		{
			TestName:      "Test Delete post OK",
			ID:            "existing_post_id",
			UserID:        "existing_user_id",
			Error:         nil,
			ExpectedError: nil,
		},
		{
			TestName:      "Test Delete post KO",
			ID:            "non_existing_post_id",
			UserID:        "non_existing_user_id",
			Error:         fmt.Errorf("error"),
			ExpectedError: ErrDeletingPost,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/posts/%v/%v", test.ID, test.UserID), strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			fakeHandler := APIHandler{
				Usecases: MockUseCases{
					Error: test.Error,
				},
			}
			if err := fakeHandler.DeletePost(c); err != test.ExpectedError {
				t.Errorf("got %v, expected %v", err, test.ExpectedError)
			}
		})
	}
}
