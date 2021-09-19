package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_TaskUpdate(t *testing.T) {
	body := strings.NewReader(`{"priority": 5}`)

	e := echo.New()
	e.Validator = &Validator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/6c0dd3e9-f7e7-478a-80cd-a1309f22f4dd", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("6c0dd3e9-f7e7-478a-80cd-a1309f22f4dd")

	h := New()
	assert.NoError(t, h.TaskUpdate(c))
	assert.Equal(t, true, strings.Contains(rec.Body.String(), "Task successfully updated"))
}
