package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/plitn/NotesApi"
)

type Auth interface {
	CreateUser(user NotesApi.User) (int, error)
	GetUser(username, password string) (NotesApi.User, error)
}
type Note interface {
	Create(userId int, note NotesApi.Note) (int, error)
	GetAll(userId int) ([]NotesApi.Note, error)
	GetById(userId, noteId int) (NotesApi.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, data NotesApi.UpdateNoteData) error
}

type Repository struct {
	Auth
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthDB(db),
		Note: NewNoteDB(db),
	}
}
