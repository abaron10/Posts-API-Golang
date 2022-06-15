package middleware

import (
	"errors"
	"fmt"
	"github.com/abaron10/Posts-API-Golang/config"
	"github.com/abaron10/Posts-API-Golang/models"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

var (
	NO_AUTH_NEEDED = []string{"login", "signup"}
)

func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuth() models.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			_, err := GetToken(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
	}
}

func GetToken(r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("x-auth-token"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Secret), nil
	})
	return token, err
}

func ValidToken(req *http.Request) (*models.AppClaims, error) {
	token, err := GetToken(req)
	if err != nil {
		return nil, errors.New("Error parsing AuthToken.")
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid credentials.")
}

func Logging() models.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("Logging auth")
			if flag {
				//aca si cumple la condicion del middleware f(w,r) ejecuta la funcion nativa, es decir la health del controller, no es una recursión es una invocación.
				f(w, r)
			} else {
				return
			}
		}
	}
}
