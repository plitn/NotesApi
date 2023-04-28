package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/plitn/NotesApi"
)

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (adb *AuthDB) CreateUser(user NotesApi.User) (int, error) {
	query := fmt.Sprintf("insert into users (name, username, password_hash) values ($1, $2, $3) returning id")
	res := adb.db.QueryRow(query, user.Name, user.Username, user.Password)
	var id int
	if err := res.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (adb *AuthDB) GetUser(username, password string) (NotesApi.User, error) {
	var user NotesApi.User
	query := fmt.Sprintf("select id from users where username=$1 and password_hash=$2")
	err := adb.db.Get(&user, query, username, password)
	return user, err
}
