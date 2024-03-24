package main

import "net/http"

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the Request (Auth header)
		// validate the token
		// get the userid from the token
		// call the handler func and continue to the endpoint
	}
}
