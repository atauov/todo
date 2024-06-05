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

func (r *TaskPostgres) CreateTask(task models.Task) (int, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (r *TaskPostgres) GetTask(taskID int) (models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, taskID).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (r *TaskPostgres) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskPostgres) UpdateTask(taskID int, input models.UpdateTask) error {
	var task models.Task
	if err := r.db.First(&task, taskID).Error; err != nil {
		return err
	}

	if input.Title != nil {
		task.Title = *input.Title
	}

	if input.Description != nil {
		task.Description = *input.Description
	}

	if err := r.db.Save(&task).Error; err != nil {
		return err
	}

	return nil
}
func (r *TaskPostgres) DeleteTask(taskID int) error {
	if err := r.db.Delete(&models.Task{}, taskID).Error; err != nil {
		return err
	}

	return nil
}
