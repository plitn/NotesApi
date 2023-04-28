package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/plitn/NotesApi/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) RoutesInit() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/log-in", h.logIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		notes := api.Group("/notes")
		{
			notes.POST("/", h.createNote)
			notes.GET("/", h.getAllNotes)
			notes.GET("/:id", h.getNote)
			notes.PUT("/:id", h.updateNote)
			notes.DELETE("/:id", h.deleteNote)
		}
	}
	return router
}
