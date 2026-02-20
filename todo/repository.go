package todo

type TodoRepository interface {
	Save(todo Todo) error
	FindAll() ([]Todo, error)
	FindByID(id string) (Todo, error)
	Delete(id string) error
}