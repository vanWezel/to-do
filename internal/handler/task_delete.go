package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/to-do/internal/model"
)

func (h *Handler) TaskDelete(c echo.Context) error {
	id := fmt.Sprintf("%v", c.Param("id"))
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("id is missing"))
	}

	if err := h.Task.Delete(id); err != nil {
		log.Print("error while saving", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "Task successfully deleted",
		Status:  true,
	})
}
