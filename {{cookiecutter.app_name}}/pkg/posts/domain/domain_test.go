package domain_test

import (
	"testing"
	"time"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain/entities"
)

func TestNewPost(t *testing.T) {
	tests := []struct {
		TestName string
		Content  string
		UserID   string
		Expected *entities.Post
	}{
		{
			"Test new post",
			"My first post",
			"1234",
			&entities.Post{
				Content: "My first post",
				Date:    time.Now(),
				UserID:  "1234",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			got := domain.NewPost(test.Content, test.UserID)
			if got.ID == "" {
				t.Errorf("Expected: %v, got: %v", test.Expected.ID, got.ID)
			}
			if got.Content != test.Expected.Content {
				t.Errorf("Expected: %v, got: %v", test.Expected.Content, got.Content)
			}
			if got.UserID != test.Expected.UserID {
				t.Errorf("Expected: %v, got: %v", test.Expected.UserID, got.UserID)
			}
		})
	}
}
