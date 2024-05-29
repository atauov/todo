package todoapp

import "time"

type TaskItem struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAT   time.Time
}
