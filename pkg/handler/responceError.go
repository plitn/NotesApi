package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type CustomError struct {
	Message string `json:"message"`
}

func newCustomError(c *gin.Context, status int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(status, message)
}
