package service

import "notes/pkg/repository"

type Services struct {
	Notes INoteService
}

func NewServices(r *repository.Repositories) *Services {
	return &Services{
		Notes: NewNoteService(r.Notes),
	}
}
