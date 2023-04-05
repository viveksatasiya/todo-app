package domain

import "time"

type ToDo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ToDoRepository interface {
	Create(toDo *ToDo) error
	FindAll() ([]ToDo, error)
	FindById(id string) (*ToDo, error)
	Update(toDo *ToDo) error
	Delete(id string) error
}
