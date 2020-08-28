package domain_test

import (
	"testing"

	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/posts/domain"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/posts/domain/entities"
)

func TestNewPost(t *testing.T) {
	tests := []struct {
		TestName string
		Content  string
		Date     string
		UserID   string
		Expected *entities.Post
	}{
		{
			"Test create post",
			"Mi first post",
			"XX/XX/XXXX",
			"",
			&entities.Post{
				Content: "Mi first post",
				Date:    "XX/XX/XXXX",
				UserID:  "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			got := domain.NewPost(test.Content, test.Date)
			if *got != *test.Expected {
				t.Errorf("Expected: %v, got: %v", test.Expected, got)
			}
		})
	}
}
