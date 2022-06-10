package controller

import (
	"encoding/json"
	"github.com/abaron10/Posts-API-Golang/errors"
	"github.com/abaron10/Posts-API-Golang/middleware"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/abaron10/Posts-API-Golang/service/post-service"
	"net/http"
)

var (
	postService post_service.PostService
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPosts(resp http.ResponseWriter, req *http.Request)
	Health(resp http.ResponseWriter, req *http.Request)
}

type postController struct{}

func NewPostController(service post_service.PostService) PostController {
	postService = service
	return &postController{}
}

func (*postController) GetPosts(resp http.ResponseWriter, req *http.Request) {
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

func (*postController) AddPosts(resp http.ResponseWriter, req *http.Request) {
	var post model.Post
	resp.Header().Set("Content-type", "application/json")
	token, err := middleware.GetToken(req)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*model.AppClaims); ok && token.Valid {
		err = json.NewDecoder(req.Body).Decode(&post)
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
		post.CreatedBy = claims.UserID
		result, err2 := postService.Create(&post)
		if err2 != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
			return
		}
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(result)
		return
	}
	resp.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(resp).Encode("Not Authorized")
}

func (*postController) Health(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode("hola")
}
