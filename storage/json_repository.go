package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/athulrajhere/todo-cli/todo"
)

type JsonRepository struct {
	filePath string
}

func (r *JsonRepository) load() ([]todo.Todo, error) {
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return []todo.Todo{}, nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read storage file: %w", err)
	}

	var todos []todo.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("failed to parse storage file: %w", err)
	}

	return todos, nil
}

func (r *JsonRepository) save(todos []todo.Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode todos: %w", err)
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write storage file: %w", err)
	}

	return nil
}

func (r *JsonRepository) Save(t todo.Todo) error {
	todos, err := r.load()
	if err != nil {
		return err
	}

	for i, existing := range todos {
		if existing.ID == t.ID {
			todos[i] = t
			return r.save(todos)
		}
	}

	return r.save(append(todos, t))
}

func (r *JsonRepository) FindAll() ([]todo.Todo, error) {
	return r.load()
}

func (r *JsonRepository) FindByID(id string) (todo.Todo, error) {
	todos, err := r.load()
	if err != nil {
		return todo.Todo{}, err
	}

	for _, t := range todos {
		if t.ID == id {
			return t, nil
		}
	}

	return todo.Todo{}, fmt.Errorf("todo with id %s not found", id)
}

func (r *JsonRepository) Delete(id string) error {
	todos, err := r.load()
	if err != nil {
		return err
	}

	for i, t := range todos {
		if t.ID == id {
			return r.save(append(todos[:i], todos[i+1:]...))
		}
	}

	return fmt.Errorf("todo with id %s not found", id)
}

func NewJsonRepository(filePath string) *JsonRepository {
	return &JsonRepository{filePath: filePath}
}