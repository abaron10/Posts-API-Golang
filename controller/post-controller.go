package controller

import (
	"encoding/json"
	"github.com/abaron10/Posts-API-Golang/errors"
	router "github.com/abaron10/Posts-API-Golang/http"
	"github.com/abaron10/Posts-API-Golang/middleware"
	"github.com/abaron10/Posts-API-Golang/models"
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
	findingResult, errFindingPosts := postService.FindAll()
	if errFindingPosts != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(findingResult)
}

func (*postController) AddPosts(resp http.ResponseWriter, req *http.Request) {
	var post models.Post
	resp.Header().Set("Content-type", "application/json")
	claims, err := middleware.ValidToken(req)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims != nil {
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
		postCreationResponse, errCreatingPost := postService.Create(&post)
		if errCreatingPost != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
			return
		}
		var postMessage = models.WebSocketMessage{
			Type:    "Post_Created",
			Payload: post,
		}
		//siq uisiera ignorar un cliente se hace aca
		router.HubS.Broadcast(postMessage, nil)
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(postCreationResponse)
		return
	}
	resp.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(resp).Encode("Not Authorized")
}

func (*postController) Health(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode("PONG")
}
