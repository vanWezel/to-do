package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/to-do/internal/model"
	"github.com/vanWezel/to-do/internal/task"
)

func (h *Handler) TaskUpdate(c echo.Context) error {
	m := task.Model{}
	if err := h.validateRequest(c, &m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := fmt.Sprintf("%v", c.Param("id"))
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("id is missing"))
	}

	if err := h.task.Update(id, &m); err != nil {
		log.Print("error while saving", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "Task successfully updated",
		Status:  true,
		Data: m,
	})
}
