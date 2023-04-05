package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viveksatasiya/todo-app/internal/app/service"
	"github.com/viveksatasiya/todo-app/internal/infrastructure/memory"

	"github.com/viveksatasiya/todo-app/cmd/api/handler"
)

func main() {
	router := mux.NewRouter()
	repo := memory.NewToDoRepository()
	service := service.NewToDoService(repo)

	handler.RegisterHandlers(router, service)

	log.Fatal(http.ListenAndServe(":8080", router))
}
