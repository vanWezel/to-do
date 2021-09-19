package task

import (
	"time"

	"github.com/vanWezel/to-do/internal/comment"
)

type Model struct {
	Id          string          `json:"id"`
	Task        string          `json:"task"`
	Priority    int             `json:"priority"`
	Labels      []string        `json:"labels"`
	Comments    []comment.Model `json:"comments"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	CompletedAt time.Time       `json:"completed_at"`
	DueAt       time.Time       `json:"due_at"`
}
