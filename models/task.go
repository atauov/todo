package models

import (
	"time"
)

type Task struct {
	ID          int        `gorm:"primary_key" json:"id"`
	Title       string     `gorm:"size:255" json:"title"`
	Description string     `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type UpdateTask struct {
	Title       *string `gorm:"size:255" json:"title"`
	Description *string `gorm:"type:varchar(255)" json:"description"`
}
