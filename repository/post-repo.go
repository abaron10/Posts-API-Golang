package repository

import (
	"github.com/abaron10/Posts-API-Golang/models"
)

type PostRepository interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}
