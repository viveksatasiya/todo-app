package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/viveksatasiya/todo-app/internal/domain"
)

func TestToDo(t *testing.T) {
	todo := domain.ToDo{
		ID:          "1",
		Title:       "Sample ToDo",
		Description: "This is a sample ToDo item",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	assert.Equal(t, "1", todo.ID)
	assert.Equal(t, "Sample ToDo", todo.Title)
	assert.Equal(t, "This is a sample ToDo item", todo.Description)
}
