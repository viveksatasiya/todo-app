package service

import (
	"github.com/viveksatasiya/todo-app/internal/domain"
)

type ToDoService struct {
	repo domain.ToDoRepository
}

func NewToDoService(repo domain.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repo}
}

func (s *ToDoService) Create(toDo *domain.ToDo) error {
	return s.repo.Create(toDo)
}

func (s *ToDoService) FindAll() ([]domain.ToDo, error) {
	return s.repo.FindAll()
}

func (s *ToDoService) FindById(id string) (*domain.ToDo, error) {
	return s.repo.FindById(id)
}

func (s *ToDoService) Update(toDo *domain.ToDo) error {
	return s.repo.Update(toDo)
}

func (s *ToDoService) Delete(id string) error {
	return s.repo.Delete(id)
}
