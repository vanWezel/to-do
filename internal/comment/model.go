package comment

import "time"

type Model struct {
	Id        string    `json:"id"`
	Message   string    `json:"message" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
