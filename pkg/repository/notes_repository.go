package repository

import "notes"

type INoteRepository interface {
	Create(input notes.Note) (*notes.Note, error)
	Get(id int) (*notes.Note, error)
	Update(input notes.Note) (*notes.Note, error)
	Delete(id int) (*notes.Note, error)
	All() ([]notes.Note, error)
}
