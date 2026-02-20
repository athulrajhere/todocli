package main

import (
	"os"

	"github.com/athulrajhere/todocli/cli"
	"github.com/athulrajhere/todocli/storage"
	"github.com/athulrajhere/todocli/todo"
)

func main() {
	repo := storage.NewJsonRepository("todos.json")
	service := todo.NewTodoService(repo)
	handler := cli.NewHandler(service)
	
	handler.Run(os.Args[1:])

}