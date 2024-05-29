package repository

import (
	"gorm.io/gorm"
	"todoapp"
)

type Task interface {
	CreateTask(task todoapp.TaskItem) (int, error)
	GetTask(taskID int) (todoapp.TaskItem, error)
	GetAllTasks() ([]todoapp.TaskItem, error)
	UpdateTask(taskID int, description string) error
	DeleteTask(taskID int) error
}

type Repository struct {
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Task: NewTaskPostgres(db),
	}
}
