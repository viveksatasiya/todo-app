package postgres_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viveksatasiya/todo-app/internal/domain"
	"github.com/viveksatasiya/todo-app/internal/infrastructure/postgres"
)

func TestToDoRepository(t *testing.T) {
	repo, err := postgres.NewToDoRepository()
	require.NoError(t, err)
	defer repo.Close()

	// Test Create
	todo := &domain.ToDo{
		Title:       "Sample ToDo",
		Description: "This is a sample ToDo item",
		CreatedAt:   time.Now(),
	}

	err = repo.Create(todo)
	require.NoError(t, err)
	require.NotZero(t, todo.ID)

	// Test FindAll
	todos, err := repo.FindAll()
	require.NoError(t, err)
	assert.Equal(t, 1, len(todos))

	// Test FindById
	found, err := repo.FindById(todo.ID)
	require.NoError(t, err)
	assert.Equal(t, todo, found)

	// Test Update
	todo.Title = "Updated ToDo"
	todo.UpdatedAt = time.Now()
	err = repo.Update(todo)
	require.NoError(t, err)

	updated, err := repo.FindById(todo.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated ToDo", updated.Title)

	// Test Delete
	err = repo.Delete(todo.ID)
	require.NoError(t, err)

	deleted, err := repo.FindById(todo.ID)
	require.Error(t, err)
	assert.Nil(t, deleted)
}
