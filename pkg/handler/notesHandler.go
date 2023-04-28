package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/plitn/NotesApi"
	"net/http"
	"strconv"
)

func (h *Handler) createNote(c *gin.Context) {
	id, ok := c.Get("UserId")
	if !ok {
		newCustomError(c, http.StatusInternalServerError, "id not found")
	}
	var data NotesApi.Note
	if err := c.BindJSON(&data); err != nil {
		newCustomError(c, http.StatusBadRequest, err.Error())
		return
	}
	noteId, err := h.services.Note.Create(id.(int), data)
	if err != nil {
		newCustomError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"noteId": noteId,
	})
}

type getAllNotesResp struct {
	Notes []NotesApi.Note `json:"data"`
}

func (h *Handler) getAllNotes(c *gin.Context) {
	id, ok := c.Get("UserId")
	if !ok {
		newCustomError(c, http.StatusInternalServerError, "id not found")
	}
	notes, err := h.services.Note.GetAll(id.(int))
	if err != nil {
		newCustomError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllNotesResp{
		Notes: notes,
	})
}
func (h *Handler) getNote(c *gin.Context) {
	id, ok := c.Get("UserId")
	if !ok {
		newCustomError(c, http.StatusInternalServerError, "id not found")
	}
	noteId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomError(c, http.StatusBadRequest, "invalid id")
		return
	}

	note, err := h.services.Note.GetById(id.(int), noteId)
	if err != nil {
		newCustomError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, note)
}
func (h *Handler) updateNote(c *gin.Context) {
	id, ok := c.Get("UserId")
	if !ok {
		newCustomError(c, http.StatusInternalServerError, "id not found")
	}
	noteId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomError(c, http.StatusBadRequest, "invalid id")
		return
	}
	var data NotesApi.UpdateNoteData
	if err := c.BindJSON(&data); err != nil {
		newCustomError(c, http.StatusBadRequest, "invalid data")
		return
	}
	err = h.services.Update(id.(int), noteId, data)
	if err != nil {
		newCustomError(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) deleteNote(c *gin.Context) {
	id, ok := c.Get("UserId")
	if !ok {
		newCustomError(c, http.StatusInternalServerError, "id not found")
	}
	noteId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newCustomError(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.Note.Delete(id.(int), noteId)
	if err != nil {
		newCustomError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
