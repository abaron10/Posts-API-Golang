package model

import "time"

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
