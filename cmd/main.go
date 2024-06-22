package main

import (
	"TODO-APP/pkg/handler"
	"TODO-APP/pkg/repository"
	"TODO-APP/pkg/service"
	"log"

	"github.com/AntonZatsepilin/todo-app"
)

func main() {
	repos := repository.NewReposytory()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
