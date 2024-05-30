package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID          int       `json:"id"`
	Title       string    `gorm:"size:255" json:"title"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAT   time.Time `json:"updated_at"`
}
