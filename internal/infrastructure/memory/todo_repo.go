package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/viveksatasiya/todo-app/internal/domain"
)

type ToDoRepository struct {
	sync.RWMutex
	items map[string]*domain.ToDo
}

func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{
		items: make(map[string]*domain.ToDo),
	}
}

func (r *ToDoRepository) Create(toDo *domain.ToDo) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.items[toDo.ID]; exists {
		return errors.New("ToDo with the given ID already exists")
	}

	toDo.CreatedAt = time.Now()
	toDo.UpdatedAt = time.Now()
	r.items[toDo.ID] = toDo
	return nil
}

func (r *ToDoRepository) FindAll() ([]domain.ToDo, error) {
	r.RLock()
	defer r.RUnlock()

	todos := make([]domain.ToDo, 0, len(r.items))
	for _, v := range r.items {
		todos = append(todos, *v)
	}
	return todos, nil
}

func (r *ToDoRepository) FindById(id string) (*domain.ToDo, error) {
	r.RLock()
	defer r.RUnlock()

	if item, exists := r.items[id]; exists {
		return item, nil
	}
	return nil, errors.New("ToDo not found")
}

func (r *ToDoRepository) Update(toDo *domain.ToDo) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.items[toDo.ID]; !exists {
		return errors.New("ToDo not found")
	}

	toDo.UpdatedAt = time.Now()
	r.items[toDo.ID] = toDo
	return nil
}

func (r *ToDoRepository) Delete(id string) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.items[id]; !exists {
		return errors.New("ToDo not found")
	}

	delete(r.items, id)
	return nil
}
