package model

import "time"

type ToDo struct {
	Id        int64
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
