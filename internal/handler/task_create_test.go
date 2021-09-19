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

func TestHandler_TaskCreate(t *testing.T) {
	body := strings.NewReader(`{"task": "Finish the test", "priority": 5, "due_at": "2021-09-19T15:04:05+02:00"}`)

	e := echo.New()
	e.Validator = &Validator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := New()
	assert.NoError(t, h.TaskCreate(c))
	assert.Equal(t, "{\"message\":\"Task successfully created\",\"status\":true,\"data\":{\"id\":\"a0bd3abc-f72b-4154-a7d5-c92ddc496a8d\",\"task\":\"Finish the test\",\"priority\":5,\"labels\":null,\"comments\":null,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"completed_at\":\"0001-01-01T00:00:00Z\",\"due_at\":\"2021-09-19T15:04:05+02:00\"}}\n", rec.Body.String())
}
