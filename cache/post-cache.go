package cache

import "github.com/abaron10/Posts-API-Golang/models"

type PostCache interface {
	Set(key string, value *models.Post)
	Get(key string) *models.Post
}
