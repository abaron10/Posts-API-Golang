package controller

import (
	"encoding/json"
	"github.com/abaron10/Posts-API-Golang/errors"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/abaron10/Posts-API-Golang/service/user-service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var (
	userService user_service.UserService
)

type UserController interface {
	SignIn(response http.ResponseWriter, request *http.Request)
	Login(resp http.ResponseWriter, req *http.Request)
}

type userController struct{}

func NewUserController(service user_service.UserService) UserController {
	userService = service
	return &userController{}
}

func (u *userController) SignIn(resp http.ResponseWriter, req *http.Request) {
	var user model.User
	resp.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := userService.ValidateUser(&user)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errHash != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errHash.Error()})
		return
	}
	user.Password = string(hashedPassword)

	result, err2 := userService.SignIn(&user)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err2.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}

func (u userController) Login(resp http.ResponseWriter, req *http.Request) {
	var user model.LoginUser
	resp.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	response, err := userService.LogIn(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(response)
}
