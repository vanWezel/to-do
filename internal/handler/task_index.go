package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/to-do/internal/model"
)

func (h *Handler) TaskIndex(c echo.Context) error {
	page := getPageQueryParam(c.QueryParam("page"))

	list, err := h.task.Index(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "Successfully loaded",
		Status:  true,
		Data:    list,
	})
}
