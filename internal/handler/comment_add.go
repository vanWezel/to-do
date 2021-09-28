package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vanWezel/to-do/internal/comment"
	"github.com/vanWezel/to-do/internal/model"
)

func (h *Handler) CommentAdd(c echo.Context) error {
	m := comment.Model{}
	if err := h.validateRequest(c, &m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := fmt.Sprintf("%v", c.Param("id"))
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("id is missing"))
	}

	if err := h.Comment.Add(id, &m); err != nil {
		log.Print("error while saving", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, model.Response{
		Message: "Comment successfully added",
		Status:  true,
		Data:    m,
	})
}
