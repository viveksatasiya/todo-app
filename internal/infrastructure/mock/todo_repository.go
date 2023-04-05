package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/viveksatasiya/todo-app/internal/domain"
)

type ToDoRepository struct {
	mock.Mock
}

func (m *ToDoRepository) Create(todo *domain.ToDo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *ToDoRepository) FindAll() ([]domain.ToDo, error) {
	args := m.Called()
	return args.Get(0).([]domain.ToDo), args.Error(1)
}

func (m *ToDoRepository) FindById(id string) (*domain.ToDo, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.ToDo), args.Error(1)
}

func (m *ToDoRepository) Update(todo *domain.ToDo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *ToDoRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
