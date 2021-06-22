package inmemory_test

import (
	"testing"
	"time"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain/entities"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/infrastructure/adapters/inmemory"
)

type mockPostAdapter struct {
	Error error
	Post  *entities.Post
}

func mockedTime() time.Time {
	return time.Date(2020, 11, 17, 20, 34, 58, 651387237, time.UTC)
}

func TestCreate(t *testing.T) {
	tests := []struct {
		TestName      string
		Post          *entities.Post
		Expected      *entities.Post
		ExpectedError error
	}{
		{
			"Test create user",
			&entities.Post{
				ID:      "1234",
				Content: "Ruben",
				Date:    mockedTime(),
				UserID:  "5678",
			},
			&entities.Post{
				ID:      "1234",
				Content: "Ruben",
				Date:    mockedTime(),
				UserID:  "5678",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			postAdapter := inmemory.PostAdapter{}
			post, err := postAdapter.Create(test.Post)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *post != *test.Expected {
				t.Errorf("Expected: %v, got: %v", test.Expected, post)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		Post          *entities.Post
		Expected      *entities.Post
		ExpectedError error
	}{
		{
			"Test get post",
			"1234",
			&entities.Post{
				ID:      "1234",
				Content: "Ruben",
				Date:    mockedTime(),
				UserID:  "5678",
			},
			&entities.Post{
				ID:      "1234",
				Content: "Ruben",
				Date:    mockedTime(),
				UserID:  "5678",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			postAdapter := inmemory.PostAdapter{
				Post: test.Post,
			}
			post, err := postAdapter.Get(test.ID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *post != *test.Expected {
				t.Errorf("Expected: %v, got: %v", test.Expected, post)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		Post          *entities.Post
		ExpectedError error
	}{
		{
			"Test delete post",
			"1234",
			&entities.Post{
				ID:      "1234",
				Content: "Ruben",
				Date:    mockedTime(),
				UserID:  "5678",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			postAdapter := inmemory.PostAdapter{
				Post: test.Post,
			}
			err := postAdapter.Delete(test.ID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
		})
	}
}
