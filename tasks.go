package main

import (
	"io"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	defer r.Body.Close()
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {
}
