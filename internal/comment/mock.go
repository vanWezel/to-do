package comment

type Mock struct {
	Interface
}

func NewMock() *Mock {
	return &Mock{}
}

func (s *Mock) Add(id string, comment *Model) error {
	comment.Id = "b2ff6329-9023-4776-a0ed-ff5fa98a888d"
	return nil
}

func (s *Mock) Delete(id string) error {
	return nil
}
