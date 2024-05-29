package handlers

import (
	"github.com/labstack/echo/v4"
	"todoapp/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}
func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()
	api := router.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.createTask)
			tasks.GET("/", h.getAllTasks)
			tasks.GET("/:id", h.getTask)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}
	}

	return router
}
