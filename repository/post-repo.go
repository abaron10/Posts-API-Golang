package repository

import (
	"github.com/abaron10/Posts-API-Golang/model"
)

type PostRepository interface {
	Save(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}
