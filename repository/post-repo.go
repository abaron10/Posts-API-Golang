package repository

import (
	"RESTapi-2/model"
)

type PostRepository interface {
	Save(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}

