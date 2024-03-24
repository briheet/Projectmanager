package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// our services

	usersService := NewUserService(s.store)
	usersService.RegisterRoutes(subrouter)

	tasksService := NewTaskService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Printf("starting our service at %s", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
