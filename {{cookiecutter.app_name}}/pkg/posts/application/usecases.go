package application

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/posts/domain/entities"
)

// Usecases represent the user business logic
type Usecases interface {
	CreatePost(content, userID string) (*entities.Post, error)
	DeletePost(id string) error
	GetPost(id string) (*entities.Post, error)
}

type postUsecase struct {
	postRepository domain.Repository
}

// NewUserUsecase will create new an userUsecase of userRepository
func NewUserUsecase(ur domain.Repository) Usecases {
	return &postUsecase{
		postRepository: ur,
	}
}

func (pu postUsecase) CreatePost(content, userID string) (*entities.Post, error) {
	post := domain.NewPost(content, userID)
	return pu.postRepository.Create(post)
}

func (pu postUsecase) DeletePost(id string) error {
	return pu.postRepository.Delete(id)
}

func (pu postUsecase) GetPost(id string) (*entities.Post, error) {
	return pu.postRepository.Get(id)
}
