package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/plitn/NotesApi"
)

type NoteDB struct {
	db *sqlx.DB
}

func NewNoteDB(db *sqlx.DB) *NoteDB {
	return &NoteDB{db: db}
}
func (n *NoteDB) Create(userId int, note NotesApi.Note) (int, error) {
	tx, err := n.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	insertNoteQuery := fmt.Sprintf("insert into notes (title, text) values ($1, $2) returning id")
	row := tx.QueryRow(insertNoteQuery, note.Title, note.Text)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	insertUserQuery := fmt.Sprintf("insert into users_notes (user_id, note_id) values ($1, $2)")
	_, err = tx.Exec(insertUserQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (n *NoteDB) GetAll(userId int) ([]NotesApi.Note, error) {
	var notes []NotesApi.Note
	getAllQuery := fmt.Sprintf("select tl.id, tl.title, tl.text from notes tl inner join users_notes ul on tl.id =ul.note_id where ul.user_id = $1 ")
	err := n.db.Select(&notes, getAllQuery, userId)
	return notes, err
}

func (n *NoteDB) GetById(userId, noteId int) (NotesApi.Note, error) {
	var note NotesApi.Note
	getNoteQuery := fmt.Sprintf("select tl.id, tl.title, tl.text from notes tl inner join users_notes ul on tl.id =ul.note_id where ul.user_id = $1 and ul.note_id=$2")
	err := n.db.Get(&note, getNoteQuery, userId, noteId)
	return note, err
}

func (n *NoteDB) Delete(userId, noteId int) error {
	query := fmt.Sprintf("delete from notes tl using users_notes ul where tl.id = ul.note_id and ul.user_id = $1 and ul.note_id = $2")
	_, err := n.db.Exec(query, userId, noteId)
	return err
}

func (n *NoteDB) Update(userId, noteId int, data NotesApi.UpdateNoteData) error {
	query := fmt.Sprintf("UPDATE notes tl SET title=$1, text=$2  FROM users_notes ul WHERE tl.id = ul.note_id AND ul.note_id=$3 AND ul.user_id=$4")
	_, err := n.db.Exec(query, data.Title, data.Text, noteId, userId)
	return err
}
