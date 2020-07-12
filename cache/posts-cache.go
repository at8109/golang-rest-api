package cache

import "github.com/at8109/golang-rest-api/entity"

type PostCache interface {
	Set(key string, value entity.Post)
	Get(key string) *entity.Post
}
