package repository

import "gorm.io/gorm"

const tasksTable = "tasks"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	//TODO
	return nil, nil
}
