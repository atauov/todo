package transport

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"todoapp/models"
)

type Handler struct {
	services *Service
}

func NewHandler(services *Service) *Handler {
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

func (h *Handler) createTask(c echo.Context) {
	var input models.Task
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateTask(models.Task{Title: input.Title, Description: input.Description})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
	})

}

func (h *Handler) getAllTasks(c echo.Context) {

}

func (h *Handler) getTask(c echo.Context) {

}

func (h *Handler) updateTask(c echo.Context) {

}

func (h *Handler) deleteTask(c echo.Context) {

}
