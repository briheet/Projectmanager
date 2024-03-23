package main

import (
	"time"
)

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectID    int64     `json:"projectid"`
	AssignedToID int64     `json:"assignedto"`
	CreatedAt    time.Time `json:"createdAt"`
}
