package main

import "database/sql"

type Store interface {
	// User
	CreateUser() error
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
	GetUserByID(id string) (*User, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)",
		t.Name, t.Status, t.ProjectID, t.AssignedToID)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow("SELECT id, name, project_id, assigned_to, createdAt FROM tasks WHERE id=?", id).Scan(&t.ID, &t.Name, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)

	return &t, err
}

func (s *Storage) GetUserByID(userID string) (*Task, error) {
	return &User{}, nil
}
