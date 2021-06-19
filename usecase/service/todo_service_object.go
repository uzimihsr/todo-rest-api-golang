package service

import "time"

// Request/Response object
type ToDoObject struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListOption struct {
	Done string
}
