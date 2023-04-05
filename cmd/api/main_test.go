package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/viveksatasiya/todo-app/cmd/api/handler"
	"github.com/viveksatasiya/todo-app/internal/app/service"
	"github.com/viveksatasiya/todo-app/internal/domain"
	"github.com/viveksatasiya/todo-app/internal/infrastructure/memory"
)

func TestAPI(t *testing.T) {
	router := mux.NewRouter()
	repo := memory.NewToDoRepository()
	svc := service.NewToDoService(repo)
	handler.RegisterHandlers(router, svc)

	// Test Create
	todo := domain.ToDo{
		ID:          "1",
		Title:       "Sample ToDo",
		Description: "This is a sample ToDo item",
	}

	body, err := json.Marshal(todo)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	// Test FindAll
	req = httptest.NewRequest("GET", "/todos", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var todos []domain.ToDo
	err = json.Unmarshal(rec.Body.Bytes(), &todos)
	require.NoError(t, err)
	assert.Equal(t, 1, len(todos))

	// Test FindById
	req = httptest.NewRequest("GET", "/todos/1", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var found domain.ToDo
	err = json.Unmarshal(rec.Body.Bytes(), &found)
	require.NoError(t, err)
	assert.Equal(t, "1", found.ID)

	// Test Update
	todo.Title = "Updated ToDo"
	body, err = json.Marshal(todo)
	require.NoError(t, err)

	req = httptest.NewRequest("PUT", "/todos/1", bytes.NewReader(body))
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)

	// Test Delete
	req = httptest.NewRequest("DELETE", "/todos/1", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}
