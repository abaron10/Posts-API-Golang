package service

import (
	"RESTapi-2/model"
	"RESTapi-2/repository"
	"errors"
	"math/rand"
)

type PostService interface {
	Validate(post *model.Post) error
	Create(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *model.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title id empty")
		return err
	}
	return nil
}

func (*service) Create(post *model.Post) (*model.Post, error) {
	post.Id = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]model.Post, error) {
	return repo.FindAll()
}
