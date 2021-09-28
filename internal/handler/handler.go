package handler

import (
	"github.com/vanWezel/to-do/internal/comment"
	"github.com/vanWezel/to-do/internal/task"
)

type Handler struct {
	Task    task.Interface
	Comment comment.Interface
}

func NewMock() *Handler {
	return &Handler{
		Task:    task.NewMock(),
		Comment: comment.NewMock(),
	}
}
