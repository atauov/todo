package services

import (
	"todoapp"
	"todoapp/internal/repository"
)

type Task interface {
	CreateTask(task todoapp.TaskItem) (int, error)
	GetTask(taskID int) (todoapp.TaskItem, error)
	GetAllTasks() ([]todoapp.TaskItem, error)
	UpdateTask(taskID int, description string) error
	DeleteTask(taskID int) error
}

type Service struct {
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Task: NewTaskService(repos.Task)}
}
