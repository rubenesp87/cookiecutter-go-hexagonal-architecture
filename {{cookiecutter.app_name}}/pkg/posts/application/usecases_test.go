package application

import (
	"errors"
	"testing"
	"time"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain/entities"
)

var errCreatingPost = errors.New("Error creating new Post")
var errGetingPost = errors.New("Error geting a Post")
var errDeletingPost = errors.New("Error deleting a Post")

type mockPostRepository struct {
	Post  *entities.Post
	Error error
}

func (m mockPostRepository) Create(post *entities.Post) (*entities.Post, error) {
	return m.Post, m.Error
}

func (m mockPostRepository) Get(id string) (*entities.Post, error) {
	return m.Post, m.Error
}

func (m mockPostRepository) Delete(id string) error {
	return m.Error
}

func TestCreatePost(t *testing.T) {
	tests := []struct {
		TestName      string
		Content       string
		UserID        string
		Post          *entities.Post
		Error         error
		ExpectedError error
	}{
		{
			"Test create post OK",
			"My first post",
			"1",
			&entities.Post{
				ID:      "1234",
				Content: "My first post",
				Date:    time.Now(),
				UserID:  "1",
			},
			nil,
			nil,
		},
		{
			"Test create post KO",
			"My wrong post",
			"",
			&entities.Post{},
			errCreatingPost,
			errCreatingPost,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {

			fakeUseCases := postUsecase{
				postRepository: mockPostRepository{
					Post:  test.Post,
					Error: test.Error,
				},
			}
			post, err := fakeUseCases.CreatePost(test.Content, test.UserID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *post != *test.Post {
				t.Errorf("Expected: %v, got: %v", test.Post, post)
			}
		})
	}
}

func TestGetPost(t *testing.T) {
	tests := []struct {
		TestName      string
		ID            string
		Post          *entities.Post
		Error         error
		ExpectedError error
	}{
		{
			"Test get user OK",
			"1234",
			&entities.Post{
				ID:      "1234",
				Content: "Ruben",
				Date:    time.Now(),
				UserID:  "1",
			},
			nil,
			nil,
		},
		{
			"Test get user KO",
			"1234",
			&entities.Post{},
			errGetingPost,
			errGetingPost,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {

			fakeUseCases := postUsecase{
				postRepository: mockPostRepository{
					Post:  test.Post,
					Error: test.Error,
				},
			}
			post, err := fakeUseCases.GetPost(test.ID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
			if *post != *test.Post {
				t.Errorf("Expected: %v, got: %v", test.Post, post)
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
			"Test delete post OK",
			"1234",
			nil,
			nil,
		},
		{
			"Test delete post KO",
			"1234",
			errDeletingPost,
			errDeletingPost,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {

			fakeUseCases := postUsecase{
				postRepository: mockPostRepository{
					Error: test.Error,
				},
			}
			err := fakeUseCases.DeletePost(test.ID)
			if err != test.ExpectedError {
				t.Errorf("Expected: %v, got: %v", test.ExpectedError, err)
			}
		})
	}
}
