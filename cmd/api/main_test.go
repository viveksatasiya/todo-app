package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/viveksatasiya/todo-app/cmd/api/handler"
	"github.com/viveksatasiya/todo-app/internal/app/service"
	"github.com/viveksatasiya/todo-app/internal/domain"
	mockRepo "github.com/viveksatasiya/todo-app/internal/infrastructure/mock"
)

func TestCreateToDo(t *testing.T) {
	repo := &mockRepo.ToDoRepository{}
	svc := service.NewToDoService(repo)
	router := mux.NewRouter()
	handler.RegisterHandlers(router, svc)

	todo := &domain.ToDo{Title: "Test task", Description: "Test task description"}

	// Set up the expectation for the Create method
	repo.On("Create", mock.MatchedBy(func(t *domain.ToDo) bool {
		return t.Title == todo.Title && t.Description == todo.Description
	})).Return(nil)

	todoJSON, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(todoJSON))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Assert the expectation was met
	repo.AssertExpectations(t)
}

func TestGetAllToDos(t *testing.T) {
	repo := &mockRepo.ToDoRepository{}
	svc := service.NewToDoService(repo)
	router := mux.NewRouter()
	handler.RegisterHandlers(router, svc)

	// Define the expected list of ToDos to be returned by the mock repository
	expectedToDos := []domain.ToDo{
		{ID: "1", Title: "Test task 1", Description: "Test task description 1"},
		{ID: "2", Title: "Test task 2", Description: "Test task description 2"},
	}

	// Set up the expectation for the FindAll method
	repo.On("FindAll").Return(expectedToDos, nil)

	req, _ := http.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todos []domain.ToDo
	if err := json.Unmarshal(rr.Body.Bytes(), &todos); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if !reflect.DeepEqual(todos, expectedToDos) {
		t.Errorf("handler returned unexpected body: got %v want %v", todos, expectedToDos)
	}

	// Assert the expectation was met
	repo.AssertExpectations(t)
}

func TestGetToDoById(t *testing.T) {
	repo := &mockRepo.ToDoRepository{}
	svc := service.NewToDoService(repo)
	router := mux.NewRouter()
	handler.RegisterHandlers(router, svc)

	// Define the expected ToDo to be returned by the mock repository
	expectedToDo := domain.ToDo{
		ID:          "1",
		Title:       "Test task",
		Description: "Test task description",
	}

	// Set up the expectation for the FindById method
	repo.On("FindById", expectedToDo.ID).Return(&expectedToDo, nil)

	req, _ := http.NewRequest("GET", "/todos/"+expectedToDo.ID, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todo domain.ToDo
	if err := json.Unmarshal(rr.Body.Bytes(), &todo); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if !reflect.DeepEqual(todo, expectedToDo) {
		t.Errorf("handler returned unexpected body: got %v want %v", todo, expectedToDo)
	}

	// Assert the expectation was met
	repo.AssertExpectations(t)
}

func TestUpdateToDo(t *testing.T) {
	repo := &mockRepo.ToDoRepository{}
	svc := service.NewToDoService(repo)
	router := mux.NewRouter()
	handler.RegisterHandlers(router, svc)

	// Define the expected updated ToDo
	updatedToDo := domain.ToDo{
		ID:          "1",
		Title:       "Updated test task",
		Description: "Updated test task description",
	}

	// Set up the expectation for the Update method
	repo.On("Update", mock.MatchedBy(func(t *domain.ToDo) bool {
		return t.ID == updatedToDo.ID && t.Title == updatedToDo.Title && t.Description == updatedToDo.Description
	})).Return(nil)

	reqBody, _ := json.Marshal(updatedToDo)
	req, _ := http.NewRequest("PUT", "/todos/"+updatedToDo.ID, bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Assert the expectation was met
	repo.AssertExpectations(t)
}

func TestDeleteToDo(t *testing.T) {
	repo := &mockRepo.ToDoRepository{}
	svc := service.NewToDoService(repo)
	router := mux.NewRouter()
	handler.RegisterHandlers(router, svc)

	// Define the ID of the ToDo to be deleted
	toDoID := "1"

	// Set up the expectation for the Delete method
	repo.On("Delete", toDoID).Return(nil)

	req, _ := http.NewRequest("DELETE", "/todos/"+toDoID, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Assert the expectation was met
	repo.AssertExpectations(t)
}
