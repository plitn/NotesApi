package service

import (
	"github.com/plitn/NotesApi"
	"github.com/plitn/NotesApi/pkg/repository"
)

type NotesService struct {
	repo repository.Note
}

func NewNotesService(repo repository.Note) *NotesService {
	return &NotesService{repo: repo}
}

func (ns *NotesService) Create(userId int, note NotesApi.Note) (int, error) {
	return ns.repo.Create(userId, note)
}

func (ns *NotesService) GetAll(userId int) ([]NotesApi.Note, error) {
	return ns.repo.GetAll(userId)
}

func (ns *NotesService) GetById(userId, noteId int) (NotesApi.Note, error) {
	return ns.repo.GetById(userId, noteId)
}

func (ns *NotesService) Delete(userId, noteId int) error {
	return ns.repo.Delete(userId, noteId)
}

func (ns *NotesService) Update(userId, noteId int, data NotesApi.UpdateNoteData) error {
	return ns.repo.Update(userId, noteId, data)
}
