package application

import (
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/posts/domain"
	"github.com/rubenesp87/cookiecutter-go-hexagonal-architecture/pkg/posts/domain/entities"
)

// Usecases represent the user business logic
type Usecases interface {
	CreatePost(content, date string) error
	DeletePost(id string) (int64, error)
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

func (pu postUsecase) CreatePost(content, date string) error {
	post := domain.NewPost(content, date)
	return pu.postRepository.Create(post)
}

func (pu postUsecase) DeletePost(id string) (int64, error) {
	return pu.postRepository.Delete(id)
}

func (pu postUsecase) GetPost(id string) (*entities.Post, error) {
	return pu.postRepository.Get(id)
}
