package main

import (
	"github.com/telega815/todo-app"
	"github.com/telega815/todo-app/pkg/handler"
	"github.com/telega815/todo-app/pkg/repository"
	"github.com/telega815/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
