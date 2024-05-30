package transport

import (
	"todoapp/internal/repository"
	"todoapp/models"
)

type Service struct {
	repository.Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Task: NewTaskService(repos.Task)}
}

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.Task) (int, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTask(taskID int) (models.Task, error) {
	return s.repo.GetTask(taskID)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTask(taskID int, description string) error {
	return s.repo.UpdateTask(taskID, description)
}

func (s *TaskService) DeleteTask(taskID int) error {
	return s.repo.DeleteTask(taskID)
}
