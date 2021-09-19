package comment

type Interface interface {
	Add(id string, comment *Model) error
	Delete(id string) error
}
