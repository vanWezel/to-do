package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)


type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err == nil {
		return nil
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
