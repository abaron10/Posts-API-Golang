package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(response http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(response http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Println("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
