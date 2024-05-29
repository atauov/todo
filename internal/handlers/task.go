package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"todoapp/models"
)

func (h *Handler) createTask(c echo.Context) {
	var input models.Input
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateTask(models.TaskItem{Title: input.Title, Description: input.Description})
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
