package controller

import (
	"encoding/json"
	"github.com/abaron10/Posts-API-Golang/errors"
	"github.com/abaron10/Posts-API-Golang/models"
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
	var user models.User
	resp.Header().Set("Content-type", "application/json")
	errDecoding := json.NewDecoder(req.Body).Decode(&user)
	if errDecoding != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	errValidation := userService.ValidateUser(&user)
	if errValidation != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errValidation.Error()})
		return
	}
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errHash != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errHash.Error()})
		return
	}
	user.Password = string(hashedPassword)
	signInResponse, errSignIn := userService.SignIn(&user)
	if errSignIn != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errSignIn.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(signInResponse)
}

func (u userController) Login(resp http.ResponseWriter, req *http.Request) {
	var user models.LoginUser
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
