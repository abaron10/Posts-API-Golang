package main

import (
	"fmt"
	"github.com/abaron10/Posts-API-Golang/config"
	"github.com/abaron10/Posts-API-Golang/controller"
	"github.com/abaron10/Posts-API-Golang/http"
	"github.com/abaron10/Posts-API-Golang/middleware"
	"github.com/abaron10/Posts-API-Golang/models"
	"github.com/abaron10/Posts-API-Golang/repository"
	"github.com/abaron10/Posts-API-Golang/repository/firestore"
	"github.com/abaron10/Posts-API-Golang/service/post-service"
	"github.com/abaron10/Posts-API-Golang/service/user-service"
	"net/http"
)

var (
	Conf                                     = config.GetConfig()
	postRepository repository.PostRepository = firestore.NewFirestorePostRepository()
	userRepository repository.UserRepository = firestore.NewFirestoreUserRepository()
	postService    post_service.PostService  = post_service.NewPostService(postRepository)
	userService    user_service.UserService  = user_service.NewUserService(userRepository)
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController(postService)
	userController controller.UserController = controller.NewUserController(userService)
)

func main() {
	httpRouter.GET("/", func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(response, "Up and running")
	})
	httpRouter.POST("/posts", AddMidleware(postController.AddPosts, middleware.CheckAuth()))
	httpRouter.GET("/posts", AddMidleware(postController.GetPosts, middleware.CheckAuth()))
	httpRouter.GET("/health", AddMidleware(postController.Health, middleware.CheckAuth(), middleware.Logging()))
	httpRouter.POST("/signin", userController.SignIn)
	httpRouter.POST("/login", userController.Login)
	httpRouter.WEBSOCKET("/ws", router.HubS.HandleWebSocket)
	httpRouter.SERVE(Conf.Port)
}

func AddMidleware(f http.HandlerFunc, middlewares ...models.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
