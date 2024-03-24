package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	errFirstNameRequired = errors.New("first name is required")
	errLastNameRequired  = errors.New("last name is required")
	errEmailRequired     = errors.New("email is required")
	errPasswordRequired  = errors.New("password is required")
)

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading the body of user during register", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user *User

	err = json.Unmarshal(body, &user)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	if err := validateUserPayload(user); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}
}

func validateUserPayload(user *User) error {
	if user.FirstName == "" {
		return errFirstNameRequired
	}

	if user.LastName == "" {
		return errLastNameRequired
	}

	if user.Email == "" {
		return errEmailRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}
