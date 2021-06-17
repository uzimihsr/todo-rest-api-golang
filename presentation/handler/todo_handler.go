package handler

import (
	"net/http"

	"github.com/uzimihsr/todo-rest-api-golang/usecase/service"
)

type ToDoHandler interface {
	Create() http.HandlerFunc
	Read() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	List() http.HandlerFunc
}

type toDoHandler struct {
	service service.ToDoService
}

func NewToDoHandler() ToDoHandler {
	return new(toDoHandler)
}

func (h *toDoHandler) Create() http.HandlerFunc {
	// WIP
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ERROR! Create ToDo API is still a work in progress.", http.StatusInternalServerError)
	}
}

func (h *toDoHandler) Read() http.HandlerFunc {
	// WIP
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ERROR! Read ToDo API is still a work in progress.", http.StatusInternalServerError)
	}
}

func (h *toDoHandler) Update() http.HandlerFunc {
	// WIP
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ERROR! Update ToDo API is still a work in progress.", http.StatusInternalServerError)
	}
}

func (h *toDoHandler) Delete() http.HandlerFunc {
	// WIP
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ERROR! Delete ToDo API is still a work in progress.", http.StatusInternalServerError)
	}
}

func (h *toDoHandler) List() http.HandlerFunc {
	// WIP
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ERROR! List ToDo API is still a work in progress.", http.StatusInternalServerError)
	}
}
