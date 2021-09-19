package handler

import (
	"github.com/vanWezel/to-do/internal/comment"
	"github.com/vanWezel/to-do/internal/task"
)

type Handler struct {
	task task.Interface
	comment comment.Interface
}

func New() *Handler {
	return &Handler{
		task: task.NewMock(),
		comment: comment.NewMock(),
	}
}
