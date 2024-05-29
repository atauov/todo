package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
)

type idResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c echo.Context, statusCode int, message string) {
	slog.Error(message)
	c.JSON(statusCode, errorResponse{message})
}
