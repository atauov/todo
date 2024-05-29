package repository

import (
	"gorm.io/gorm"
	"todoapp"
)

type TaskPostgres struct {
	db *gorm.DB
}

func NewTaskPostgres(db *gorm.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) CreateTask(task todoapp.TaskItem) (int, error) {
	return 0, nil
}
func (r *TaskPostgres) GetTask(taskID int) (todoapp.TaskItem, error) {
	return todoapp.TaskItem{}, nil
}
func (r *TaskPostgres) GetAllTasks() ([]todoapp.TaskItem, error) {
	return nil, nil
}
func (r *TaskPostgres) UpdateTask(taskID int, description string) error {
	return nil
}
func (r *TaskPostgres) DeleteTask(taskID int) error {
	return nil
}
