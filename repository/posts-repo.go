package repository

import "github.com/at8109/golang-rest-api/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(postID string) (*entity.Post, error)
	DeleteByID(postID string) error
	UpdateByID(postID string) error
}
