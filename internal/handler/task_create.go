package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/to-do/internal/model"
	"github.com/vanWezel/to-do/internal/task"
)

func (h *Handler) TaskCreate(c echo.Context) error {
	m := task.Model{}
	if err := h.validateRequest(c, &m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.Task.Create(&m); err != nil {
		log.Print("error while saving", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, model.Response{
		Message: "Task successfully created",
		Status:  true,
		Data:    m,
	})
}
