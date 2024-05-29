package services

import (
	"todoapp"
	"todoapp/internal/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task todoapp.TaskItem) (int, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTask(taskID int) (todoapp.TaskItem, error) {
	return s.repo.GetTask(taskID)
}

func (s *TaskService) GetAllTasks() ([]todoapp.TaskItem, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTask(taskID int, description string) error {
	return s.repo.UpdateTask(taskID, description)
}

func (s *TaskService) DeleteTask(taskID int) error {
	return s.repo.DeleteTask(taskID)
}
