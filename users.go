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

	if err = validateUserPayload(user); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	// it is important to store userpassword by hashing it and not sending the real one to database

	hashedPW, err := HashPassword(user.Password)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "problem in hashing the password"})
		return
	}

	user.Password = hashedPW

	u, err := s.store.CreateUser(user)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating a user"})
		return
	}

	token, err := createAndSetAuthCookies(u.ID, w)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating session"})
		return
	}

	WriteJson(w, http.StatusCreated, token)
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

func createAndSetAuthCookies(userid int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, userid)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
