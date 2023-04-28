package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newCustomError(c, http.StatusUnauthorized, "auth empty header")
		return
	}
	headerWods := strings.Split(header, " ")
	if len(headerWods) != 2 {
		newCustomError(c, http.StatusUnauthorized, "auth strange header")
		return
	}
	userId, err := h.services.ParseToken(headerWods[1])
	if err != nil {
		newCustomError(c, http.StatusUnauthorized, "parse err")
		return
	}
	c.Set("UserId", userId)
}
