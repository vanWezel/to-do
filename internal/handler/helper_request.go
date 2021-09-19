package handler

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) validateRequest(c echo.Context, m interface{}) error {
	if err := c.Bind(m); err != nil {
		log.Println("error while binding", err)
		return err
	}

	if err := c.Validate(m); err != nil {
		log.Print("error while validating", err)
		return err
	}

	return nil
}

func getPageQueryParam(value string) int {
	if value == "" {
		return 1
	}

	page, err := strconv.Atoi(value)
	if err != nil {
		log.Println("error while parsing page param", err)
	}

	if page == 0 {
		page = 1
	}

	return page
}
