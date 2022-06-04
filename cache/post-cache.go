package cache

import "github.com/abaron10/Posts-API-Golang/model"

type PostCache interface {
	Set(key string, value *model.Post)
	Get(key string) *model.Post
}
