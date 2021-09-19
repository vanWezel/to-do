package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CommentDelete(t *testing.T) {
	e := echo.New()
	e.Validator = &Validator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, "/e3dc06f0-2568-42de-8760-7481d7408bd4/comments/b2ff6329-9023-4776-a0ed-ff5fa98a888d", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:task_id/comments/:id")
	c.SetParamNames("task_id")
	c.SetParamValues("e3dc06f0-2568-42de-8760-7481d7408bd4")
	c.SetParamNames("id")
	c.SetParamValues("b2ff6329-9023-4776-a0ed-ff5fa98a888d")

	h := New()
	assert.NoError(t, h.CommentDelete(c))
}
