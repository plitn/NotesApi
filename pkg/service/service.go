package service

import (
	"github.com/plitn/NotesApi"
	"github.com/plitn/NotesApi/pkg/repository"
)

type Auth interface {
	CreateUser(user NotesApi.User) (int, error)
	GenToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Note interface {
	Create(userId int, note NotesApi.Note) (int, error)
	GetAll(userId int) ([]NotesApi.Note, error)
	GetById(userId, noteId int) (NotesApi.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, data NotesApi.UpdateNoteData) error
}

type Service struct {
	Auth
	Note
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
		Note: NewNotesService(repo.Note),
	}
}
