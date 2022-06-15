package post_service

import (
	"errors"
	"github.com/abaron10/Posts-API-Golang/models"
	"github.com/abaron10/Posts-API-Golang/repository"
	"github.com/google/uuid"
)

type PostService interface {
	Validate(post *models.Post) error
	Create(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}

type postService struct{}

var (
	postRepository repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	postRepository = repository
	return &postService{}
}

func (*postService) Validate(post *models.Post) error {
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

func (*postService) Create(post *models.Post) (*models.Post, error) {
	post.Id = uuid.New().String()
	return postRepository.Save(post)
}

func (*postService) FindAll() ([]models.Post, error) {
	return postRepository.FindAll()
}
