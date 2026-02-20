package todo

import (
	"fmt"
	"strings"
)

type TodoService interface {
	Add(title string) (Todo, error)
	Complete(id string) error
	Delete(id string) error
	List() ([]Todo, error)
}

type todoService struct {
	repo TodoRepository
}

func (s *todoService) Add(title string) (Todo, error) {
	todo := NewTodo(title)
	err := s.repo.Save(todo)
	return todo, err
}

func (s *todoService) Complete(id string) error {
	todo, err := s.findByPrefix(id)
	if err != nil {
		return err
	}
	todo.Completed = true
	return s.repo.Save(todo)
}

func (s *todoService) Delete(id string) error {
	todo, err := s.findByPrefix(id)
    if err != nil {
        return err
    }

	return s.repo.Delete(todo.ID)
}

func (s *todoService) List() ([]Todo, error) {
	return s.repo.FindAll()
}

func (s *todoService) findByPrefix(prefix string) (Todo, error) {
    todos, err := s.repo.FindAll()
    if err != nil {
        return Todo{}, err
    }
    for _, t := range todos {
        if strings.HasPrefix(t.ID, prefix) {
            return t, nil
        }
    }
    return Todo{}, fmt.Errorf("no todo found with id starting with %s", prefix)
}


func NewTodoService(repo TodoRepository)  TodoService {
	return &todoService{repo: repo}
}