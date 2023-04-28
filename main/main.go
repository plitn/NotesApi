package main

import (
	_ "github.com/lib/pq"
	"github.com/plitn/NotesApi"
	"github.com/plitn/NotesApi/pkg/handler"
	"github.com/plitn/NotesApi/pkg/repository"
	"github.com/plitn/NotesApi/pkg/service"
	"log"
)

func main() {
	db, err := repository.NewPostgresql(repository.ConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgresql",
		Password: "admin",
		DBName:   "notes-db",
		SSlMode:  "disable",
	})
	if err != nil {
		log.Fatalf("db err: %s", err)
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)

	h := handler.NewHandler(services)
	server := new(NotesApi.Server)
	if err := server.Run("8000", h.RoutesInit()); err != nil {
		log.Fatalf("running error: %s", err)
	}
}
