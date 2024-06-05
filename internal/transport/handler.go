package transport

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
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

func (h *Handler) createTask(c echo.Context) error {
	var input models.Task
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return err
	}

	id, err := h.services.CreateTask(models.Task{Title: input.Title, Description: input.Description})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, idResponse{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
	})
}

func (h *Handler) getAllTasks(c echo.Context) error {
	tasks, err := h.services.GetAllTasks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) getTask(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id should be numeric")
		slog.Error(err.Error())
		return err
	}

	if id <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "id should be positive")
		slog.Error("id should be positive")
		return errors.New("id should be positive")
	}

	task, err := h.services.GetTask(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (h *Handler) updateTask(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id should be numeric")
		slog.Error(err.Error())
		return err
	}

	if id <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "id should be positive")
		slog.Error("id should be positive")
		return errors.New("id should be positive")
	}

	input := models.UpdateTask{}
	if err = c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return err
	}

	if err = h.services.UpdateTask(id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		slog.Error(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, "OK")
}

func (h *Handler) deleteTask(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id should be numeric")
		slog.Error(err.Error())
		return err
	}

	if id <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "id should be positive")
		slog.Error("id should be positive")
		return errors.New("id should be positive")
	}

	if err = h.services.DeleteTask(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		slog.Error(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, "OK")
}
