package post_service

import (
	"errors"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/abaron10/Posts-API-Golang/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *model.Post) error
	Create(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}

type postService struct{}

var (
	postRepository repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	postRepository = repository
	return &postService{}
}

func (*postService) Validate(post *model.Post) error {
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

func (*postService) Create(post *model.Post) (*model.Post, error) {
	post.Id = rand.Int63()
	return postRepository.Save(post)
}

func (*postService) FindAll() ([]model.Post, error) {
	return postRepository.FindAll()
}
