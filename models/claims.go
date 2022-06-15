package models

import (
	"github.com/golang-jwt/jwt"
)

type AppClaims struct {
	UserID string `json:"user_id"`
	//herencia por composición aca Appclaims tendria todo lo de standar claims
	jwt.StandardClaims
}
