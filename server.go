package main

import (
	"RESTapi-2/controller"
	"RESTapi-2/http"
	"RESTapi-2/repository"
	"RESTapi-2/service"
	"fmt"
	"net/http"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewController(postService)
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(response, "Up and running")
	})
	httpRouter.POST("/posts", postController.AddPosts)
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.SERVE(port)
}
