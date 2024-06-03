package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          int    `json:"id"`
	Title       string `gorm:"size:255" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
