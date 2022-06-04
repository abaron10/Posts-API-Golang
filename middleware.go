package main

import (
	"fmt"
	"github.com/abaron10/Posts-API-Golang/model"
	"net/http"
)

func CheckAuth() model.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("checking auth")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}
