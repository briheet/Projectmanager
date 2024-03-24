package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the Request (Auth header)
		tokenString := GetTokenFromRequest(r)
		// validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Println("failed to authenticate token")
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("failed to authenticate token")
			permissionDenied(w)
			return
		}

		// get the userid from the tokenString

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)

		_, err = store.GetUserByID(userID)
		if err != nil {
			log.Println("failed to get userID")
			permissionDenied(w)
			return
		}
		// call the handler func and continue to the endpoint

		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: fmt.Errorf("permission denied").Error()})
	return
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	tokenQuery := r.URL.Query().Get("token") //    /api/url?query=token

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func validateJWT(t string) (*jwt.Token, error) {
	secret := Envs.JWTSecret

	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
}
