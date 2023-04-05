package domain

type ToDoServiceInterface interface {
	Create(todo *ToDo) error
	FindAll() ([]ToDo, error)
	FindById(id string) (*ToDo, error)
	Update(todo *ToDo) error
	Delete(id string) error
}
