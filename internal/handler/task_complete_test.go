package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_TaskComplete(t *testing.T) {
	e := echo.New()
	e.Validator = &Validator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/e3dc06f0-2568-42de-8760-7481d7408bd4/complete", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id/complete")
	c.SetParamNames("id")
	c.SetParamValues("e3dc06f0-2568-42de-8760-7481d7408bd4")

	h := New()
	assert.NoError(t, h.TaskComplete(c))
}
