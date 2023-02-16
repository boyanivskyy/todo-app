package main

import (
	"log"

	"github.com/boyanivskyy/todo-app"
	"github.com/boyanivskyy/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	server := new(todo.Server)

	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured when running http server: %s", err.Error())
	}
}
