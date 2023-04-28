package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/plitn/NotesApi"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var data NotesApi.User
	if err := c.BindJSON(&data); err != nil {
		newCustomError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Auth.CreateUser(data)
	if err != nil {
		newCustomError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type logInData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) logIn(c *gin.Context) {

	var data logInData
	if err := c.BindJSON(&data); err != nil {
		newCustomError(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Auth.GenToken(data.Username, data.Password)
	if err != nil {
		newCustomError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
