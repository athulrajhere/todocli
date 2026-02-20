package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID string
	Title string
	Completed bool
	CreatedAt time.Time
}

func NewTodo(title string) Todo {
	return Todo{
		ID: fmt.Sprintf("%d", time.Now().UnixNano()),
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

