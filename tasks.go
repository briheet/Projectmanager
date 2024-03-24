package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	errNameRequired      = errors.New("name is required")
	errProjectIDRequired = errors.New("project id is required")
	errUserIDRequired    = errors.New("user id is required")
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
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid Request Payload!"})
		return
	}

	defer r.Body.Close()

	var task *Task

	err = json.Unmarshal(body, &task)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid Request Payload!"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}

	WriteJson(w, http.StatusCreated, t)
}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedToID == 0 {
		return errUserIDRequired
	}

	return nil
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "id is required"})
		return
	}

	t, err := s.store.GetTask(id)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "task not found"})
	}

	WriteJson(w, http.StatusOK, t)
}
