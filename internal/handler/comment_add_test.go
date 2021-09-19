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

func TestHandler_CommentAddSuccess(t *testing.T) {
	body := strings.NewReader(`{"message": "See you soon!"}`)

	e := echo.New()
	e.Validator = &Validator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/6c0dd3e9-f7e7-478a-80cd-a1309f22f4dd/comments", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id/comments")
	c.SetParamNames("id")
	c.SetParamValues("6c0dd3e9-f7e7-478a-80cd-a1309f22f4dd")

	h := New()
	assert.NoError(t, h.CommentAdd(c))
	assert.Equal(t, "{\"message\":\"Comment successfully added\",\"status\":true,\"data\":{\"id\":\"b2ff6329-9023-4776-a0ed-ff5fa98a888d\",\"message\":\"See you soon!\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"}}\n", rec.Body.String())
}

func TestHandler_CommentAddNoBody(t *testing.T) {
	e := echo.New()
	e.Validator = &Validator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/6c0dd3e9-f7e7-478a-80cd-a1309f22f4dd/comments", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id/comments")
	c.SetParamNames("id")
	c.SetParamValues("6c0dd3e9-f7e7-478a-80cd-a1309f22f4dd")

	h := New()
	assert.Error(t, h.CommentAdd(c))
}
