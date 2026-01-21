package model

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	Status      string
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
