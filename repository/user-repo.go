package repository

import "github.com/abaron10/Posts-API-Golang/model"

type UserRepository interface {
	SignIn(user *model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
