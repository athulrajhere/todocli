package main

import (
	"os"

	"github.com/athulrajhere/todo-cli/cli"
	"github.com/athulrajhere/todo-cli/storage"
	"github.com/athulrajhere/todo-cli/todo"
)

func main() {
	repo := storage.NewJsonRepository("todos.json")
	service := todo.NewTodoService(repo)
	handler := cli.NewHandler(service)
	
	handler.Run(os.Args[1:])

}