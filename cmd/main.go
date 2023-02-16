package main

import (
	"log"

	"github.com/boyanivskyy/todo-app"
	"github.com/boyanivskyy/todo-app/pkg/handler"
	"github.com/boyanivskyy/todo-app/pkg/repository"
	"github.com/boyanivskyy/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured when running http server: %s", err.Error())
	}
}
