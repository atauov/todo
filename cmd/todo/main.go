package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	todo "todoapp"
	"todoapp/internal/config"
	"todoapp/internal/repository"
	"todoapp/internal/transport"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))
	log.Info("initializing server", slog.String("address", cfg.Address))
	log.Debug("logger debug mod enabled")

	if err := godotenv.Load(); err != nil {
		log.Info("Cant load env variable: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Error("Failed to initialize database: %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := transport.NewService(repos)
	handlers := transport.NewHandler(services)

	srv := new(todo.Server)

	go func() {
		if err = srv.Run(cfg, handlers.InitRoutes()); err != nil {
			log.Error("Cant start server: %s", err.Error())
			return
		}
	}()

	log.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Shutting down server...")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Error("Server shutdown failed: %s", err.Error())
	}

	if err = repository.CloseRepository(db); err != nil {
		log.Error("Failed to close repository: %s", err.Error())
		return
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
