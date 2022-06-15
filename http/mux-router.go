package router

import (
	"fmt"
	"github.com/abaron10/Posts-API-Golang/websocket"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type muxRouter struct {
}

var (
	muxDispatcher = mux.NewRouter()
	HubS          *websocket.Hub
)

func NewMuxRouter() Router {
	HubS = websocket.NewHub()

	go HubS.Run()
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(response http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(response http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) WEBSOCKET(uri string, f func(response http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f)
}

func (*muxRouter) SERVE(port string) {
	handler := cors.AllowAll().Handler(muxDispatcher)
	fmt.Println("Mux HTTP server running on port %v", port)
	http.ListenAndServe(":"+port, handler)
}
