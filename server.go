package main

import (
	"fmt"
	"github.com/abaron10/Posts-API-Golang/controller"
	"github.com/abaron10/Posts-API-Golang/http"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/abaron10/Posts-API-Golang/repository"
	"github.com/abaron10/Posts-API-Golang/service"
	"net/http"
	"os"
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
	httpRouter.GET("/health", AddMidleware(postController.Health, CheckAuth()))

	httpRouter.SERVE(os.Getenv("PORT"))
}

func AddMidleware(f http.HandlerFunc, middlewares ...model.Middleware) http.HandlerFunc {
	fmt.Println(middlewares)
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
