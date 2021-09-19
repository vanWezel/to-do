package task

import (
	"time"
)

type Mock struct {
	Interface
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) Get(id string) (*Model, error) {
	return &Model{
		Id:          id,
		Task:        "Task one",
		Priority:    5,
		Labels:      []string{"Project One"},
		Comments:    nil,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		CompletedAt: time.Time{},
		DueAt:       time.Now().Add(7 * 24 * 60 * time.Hour).UTC(),
	}, nil
}

func (m *Mock) Index(page int) ([]Model, error) {
	var list []Model

	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		return nil, err
	}

	list = append(list, Model{
		Id:          "6c0dd3e9-f7e7-478a-80cd-a1309f22f4dd",
		Task:        "Preparing for second interview",
		Priority:    5,
		Labels:      []string{"Sam", "Georg"},
		Comments:    nil,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		CompletedAt: time.Time{},
		DueAt:       time.Date(2021, 9, 22, 13, 00, 00, 00, loc),
	})

	list = append(list, Model{
		Id:          "f1ea58a7-19f2-4ed2-a434-f4426a4b4e13",
		Task:        "Getting groceries",
		Priority:    5,
		Labels:      []string{"AH", "Jumbo"},
		Comments:    nil,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		CompletedAt: time.Time{},
		DueAt:       time.Now().Add(24 * 60 * time.Hour).UTC(),
	})

	list = append(list, Model{
		Id:          "fa62c6c2-0679-4704-97d1-998cf54e58bb",
		Task:        "Make diner",
		Priority:    2,
		Labels:      []string{"Parents", "Party"},
		Comments:    nil,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		CompletedAt: time.Time{},
		DueAt:       time.Now().Add(3 * 24 * 60 * time.Hour).UTC(),
	})

	return list, nil
}

func (m *Mock) Create(task *Model) error {
	task.Id = "a0bd3abc-f72b-4154-a7d5-c92ddc496a8d"
	return nil
}

func (m *Mock) Update(id string, task *Model) error {
	task.UpdatedAt = time.Now().UTC()
	return nil
}

func (m *Mock) Delete(id string) error {
	return nil
}

func (m *Mock) Complete(id string) error {
	return nil
}
