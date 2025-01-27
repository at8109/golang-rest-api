package service

import (
	"errors"

	"github.com/at8109/golang-rest-api/entity"
	"github.com/at8109/golang-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(postID string) (*entity.Post, error)
	DeleteByID(postID string) error
	UpdateByID(postID string) error
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The Post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The Post title is empty")
		return err
	}
	return nil
}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	return repo.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (*service) FindByID(postID string) (*entity.Post, error) {
	return repo.FindByID(postID)
}

func (*service) DeleteByID(postID string) error {
	if postID == "" {
		err := errors.New("The Post is empty")
		return err
	}
	return repo.DeleteByID(postID)
}

func (*service) UpdateByID(postID string) error {
	if postID == "" {
		err := errors.New("The Post is empty")
		return err
	}
	return repo.UpdateByID(postID)
}
