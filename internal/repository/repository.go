package repository

import (
	"gorm.io/gorm"
	"todoapp/models"
)

type Task interface {
	CreateTask(task models.Task) (int, error)
	GetTask(taskID int) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
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

func CloseRepository(db *gorm.DB) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDb.Close()
}
