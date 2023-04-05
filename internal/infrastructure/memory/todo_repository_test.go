package memory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viveksatasiya/todo-app/internal/domain"
	"github.com/viveksatasiya/todo-app/internal/infrastructure/memory"
)

func TestToDoRepository(t *testing.T) {
	repo := memory.NewToDoRepository()

	// Test Create
	todo := &domain.ToDo{
		ID:          "1",
		Title:       "Sample ToDo",
		Description: "This is a sample ToDo item",
	}

	err := repo.Create(todo)
	require.NoError(t, err)

	// Test FindAll
	todos, err := repo.FindAll()
	require.NoError(t, err)
	assert.Equal(t, 1, len(todos))

	// Test FindById
	found, err := repo.FindById("1")
	require.NoError(t, err)
	assert.Equal(t, todo, found)

	// Test Update
	todo.Title = "Updated ToDo"
	err = repo.Update(todo)
	require.NoError(t, err)

	updated, err := repo.FindById("1")
	require.NoError(t, err)
	assert.Equal(t, "Updated ToDo", updated.Title)

	// Test Delete
	err = repo.Delete("1")
	require.NoError(t, err)

	deleted, err := repo.FindById("1")
	require.Error(t, err)
	assert.Nil(t, deleted)
}
