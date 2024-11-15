package repository

type Repositories struct {
	Notes INoteRepository
}

func NewRepositories(notes INoteRepository) *Repositories {
	return &Repositories{
		Notes: notes,
	}
}
