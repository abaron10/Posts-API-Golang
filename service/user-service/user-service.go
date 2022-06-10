package user_service

import (
	"errors"
	"github.com/abaron10/Posts-API-Golang/config"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/abaron10/Posts-API-Golang/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

type UserService interface {
	SignIn(user *model.User) (*model.User, error)
	LogIn(user *model.LoginUser) (*model.LoginResponse, error)
	ValidateUser(post *model.User) error
}

type userService struct{}

var (
	userRepository repository.UserRepository
)

func NewUserService(repository repository.UserRepository) UserService {
	userRepository = repository
	return &userService{}
}

func (*userService) ValidateUser(user *model.User) error {
	if user == nil {
		err := errors.New("The user is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("The Name is empty")
		return err
	}
	if user.LastName == "" {
		err := errors.New("The LastName is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("The Email is empty")
		return err
	}
	if user.Password == "" {
		err := errors.New("The Password is empty")
		return err
	}
	if user.UserName == "" {
		err := errors.New("The UserName is empty")
		return err
	}
	foundUser, err := userRepository.GetUserByEmail(user.Email)
	if err != nil {
		err := errors.New("An error has ocurred. Please try again! ")
		return err
	}
	if foundUser.Email != "" {
		err := errors.New("There is an existing user with this email.Try with a new one.")
		return err
	}
	return nil
}

func (u *userService) SignIn(user *model.User) (*model.User, error) {
	user.Id = rand.Int63()
	return userRepository.SignIn(user)
}

func (u userService) LogIn(user *model.LoginUser) (*model.LoginResponse, error) {
	userFound, err := userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if userFound.Id == 0 {
		return nil, errors.New("Invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("Invalid credentials")
	}
	claims := model.AppClaims{UserID: strconv.FormatInt(userFound.Id, 10),
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.Secret))
	if err != nil {
		return nil, err
	}
	return &model.LoginResponse{userFound.UserName, tokenString}, nil
}
