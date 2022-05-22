package controller

import (
	"encoding/json"
	"github.com/abaron10/Posts-API-Golang/errors"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/abaron10/Posts-API-Golang/service"
	"net/http"
)

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPosts(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}

func (*controller) AddPosts(resp http.ResponseWriter, req *http.Request) {
	var post model.Post
	resp.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)

}
