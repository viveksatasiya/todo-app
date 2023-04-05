package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/viveksatasiya/todo-app/internal/app/service"
	"github.com/viveksatasiya/todo-app/internal/domain"
	mymock "github.com/viveksatasiya/todo-app/internal/infrastructure/mock"
)

func TestToDoService(t *testing.T) {
	repo := new(mymock.ToDoRepository)
	service := service.NewToDoService(repo)

	todo := &domain.ToDo{
		ID:          "1",
		Title:       "Sample ToDo",
		Description: "A sample ToDo item",
	}

	// Mock the Create method
	repo.On("Create", mock.AnythingOfType("*domain.ToDo")).Return(nil)

	// Test Create
	err := service.Create(todo)
	assert.NoError(t, err)

	// Mock the FindAll method
	repo.On("FindAll").Return([]domain.ToDo{*todo}, nil)

	// Test FindAll
	todos, err := service.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))

	// Mock the FindById method
	repo.On("FindById", todo.ID).Return(todo, nil)

	// Test FindById
	foundTodo, err := service.FindById(todo.ID)
	assert.NoError(t, err)
	assert.Equal(t, todo.Title, foundTodo.Title)

	// Mock the Update method
	repo.On("Update", mock.AnythingOfType("*domain.ToDo")).Return(nil)

	// Test Update
	todo.Title = "Updated ToDo"
	err = service.Update(todo)
	assert.NoError(t, err)

	// Mock the Delete method
	repo.On("Delete", todo.ID).Return(nil)

	// Test Delete
	err = service.Delete(todo.ID)
	assert.NoError(t, err)

	// Assert that the mocked methods were called
	repo.AssertExpectations(t)
}
