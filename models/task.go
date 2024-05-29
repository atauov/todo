package models

import (
	"gorm.io/gorm"
	"time"
)

type TaskItem struct {
	gorm.Model
	ID          int
	Title       string `gorm:"size:255"`
	Description string `gorm:"type:varchar(100)"`
	CreatedAt   time.Time
	UpdatedAT   time.Time
}
