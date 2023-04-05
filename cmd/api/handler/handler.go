package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viveksatasiya/todo-app/internal/domain"
)

func RegisterHandlers(router *mux.Router, service domain.ToDoServiceInterface) {
	router.HandleFunc("/todos", createToDo(service)).Methods("POST")
	router.HandleFunc("/todos", getAllToDos(service)).Methods("GET")
	router.HandleFunc("/todos/{id}", getToDoById(service)).Methods("GET")
	router.HandleFunc("/todos/{id}", updateToDo(service)).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteToDo(service)).Methods("DELETE")
}

// Move the createToDo, getAllToDos, getToDoById, updateToDo, and deleteToDo functions from the cmd/api/main.go file to this file.
func createToDo(service domain.ToDoServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo domain.ToDo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := service.Create(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func getAllToDos(service domain.ToDoServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := service.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(todos)
	}
}

func getToDoById(service domain.ToDoServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		todo, err := service.FindById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(todo)
	}
}

func updateToDo(service domain.ToDoServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo domain.ToDo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id := mux.Vars(r)["id"]
		todo.ID = id

		if err := service.Update(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteToDo(service domain.ToDoServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		if err := service.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
