package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/vanWezel/to-do/internal/comment"
	"github.com/vanWezel/to-do/internal/handler"
	"github.com/vanWezel/to-do/internal/task"
	"github.com/vanWezel/to-do/pkg/helper"
)

func main() {
	port := helper.Getenv("PORT", "80")

	e := echo.New()
	e.Validator = &handler.Validator{Validator: validator.New()}
	h := &handler.Handler{
		Task:    task.NewMock(),
		Comment: comment.NewMock(),
	}

	e.GET("/health", h.Health)

	// basic crud actions
	e.GET("/", h.TaskIndex)
	e.POST("/", h.TaskCreate)
	e.PUT("/:id", h.TaskUpdate)
	e.DELETE("/:id", h.TaskDelete)

	// additional task actions
	e.PUT("/:id/complete", h.TaskComplete)

	// commenting actions
	e.POST("/:task_id/comments", h.CommentAdd)
	e.DELETE("/:task_id/comments/:id", h.CommentDelete)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
