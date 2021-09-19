package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/to-do/internal/model"
)

func (h *Handler) CommentDelete(c echo.Context) error {
	id := fmt.Sprintf("%v", c.Param("id"))
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("id is missing"))
	}

	if err := h.comment.Delete(id); err != nil {
		log.Print("error while saving", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, model.Response{
		Message: "Comment successfully deleted",
		Status:  true,
	})
}
