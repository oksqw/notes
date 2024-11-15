package service

import (
	"notes"
	"notes/pkg/repository"
	"notes/pkg/request"
	"notes/pkg/validator"
)

type INoteService interface {
	Create(input request.CreateNoteRequest) (*notes.Note, error)
	Get(id int) (*notes.Note, error)
	Update(input request.UpdateNoteRequest) (*notes.Note, error)
	Delete(id int) (*notes.Note, error)
	All() ([]notes.Note, error)
}

type NoteService struct {
	repo repository.INoteRepository
}

func (service *NoteService) Get(id int) (*notes.Note, error) {
	err := validator.ValidateId(id)
	if err != nil {
		return nil, err
	}
	return service.repo.Get(id)
}

func (service *NoteService) Create(input request.CreateNoteRequest) (*notes.Note, error) {
	err := validator.ValidateCreateNoteRequest(input)
	if err != nil {
		return nil, err
	}
	return service.repo.Create(notes.Note{Title: input.Title, Text: input.Text})
}

func (service *NoteService) Update(input request.UpdateNoteRequest) (*notes.Note, error) {
	err := validator.ValidateUpdateNoteRequest(input)
	if err != nil {
		return nil, err
	}
	return service.repo.Update(notes.Note{Id: input.Id, Title: input.Title, Text: input.Text})
}

func (service *NoteService) Delete(id int) (*notes.Note, error) {
	err := validator.ValidateId(id)
	if err != nil {
		return nil, err
	}
	return service.repo.Delete(id)
}

func (service *NoteService) All() ([]notes.Note, error) {
	return service.repo.All()
}

func NewNoteService(repo repository.INoteRepository) *NoteService {
	return &NoteService{repo: repo}
}
