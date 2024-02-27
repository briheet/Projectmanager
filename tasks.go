package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	store Store
}

func NewTaskService(s Store) *TasksService {
	return &TasksService{
		store: s,
	}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTasks).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

func (s *TasksService) handleCreateTasks(w http.ResponseWriter, r *http.Request) {
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {
}
