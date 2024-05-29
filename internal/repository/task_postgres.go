package repository

import (
	"gorm.io/gorm"
	"todoapp/models"
)

type TaskPostgres struct {
	db *gorm.DB
}

func NewTaskPostgres(db *gorm.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) CreateTask(task models.TaskItem) (int, error) {
	return 0, nil
}
func (r *TaskPostgres) GetTask(taskID int) (models.TaskItem, error) {
	return models.TaskItem{}, nil
}
func (r *TaskPostgres) GetAllTasks() ([]models.TaskItem, error) {
	return nil, nil
}
func (r *TaskPostgres) UpdateTask(taskID int, description string) error {
	return nil
}
func (r *TaskPostgres) DeleteTask(taskID int) error {
	return nil
}
