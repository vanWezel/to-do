package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health")

	h := &Handler{
		task:    nil,
		comment: nil,
	}

	if assert.NoError(t, h.Health(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Yup, still working!", rec.Body.String())
	}
}
