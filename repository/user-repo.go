package repository

import "github.com/abaron10/Posts-API-Golang/models"

type UserRepository interface {
	SignIn(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
