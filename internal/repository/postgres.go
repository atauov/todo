package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todoapp/internal/config"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
