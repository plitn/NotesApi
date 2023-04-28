package NotesApi

type Note struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Text  string `json:"text" db:"text"`
}

type UsersNotes struct {
	Id     int
	UserId int
	NoteId int
}

type UpdateNoteData struct {
	Title *string `json:"title"`
	Text  *string `json:"text"`
}
