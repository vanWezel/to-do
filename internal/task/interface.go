package task

type Interface interface {
	Get(id string) (*Model, error)
	Index(page int) ([]Model, error)
	Create(*Model) error
	Update(id string, task *Model) error
	Delete(id string) error
	Complete(id string) error
}
