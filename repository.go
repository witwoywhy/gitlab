package gitlab

type Repository struct {
	Id   int
	Name string
}

func NewRepository(id int, name string) *Repository {
	return &Repository{
		Id:   id,
		Name: name,
	}
}
