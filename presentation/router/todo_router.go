package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uzimihsr/todo-rest-api-golang/presentation/handler"
)

type ToDoRouter struct {
	handler handler.ToDoHandler
	router  *mux.Router
}

func NewToDoRouter(h handler.ToDoHandler) *ToDoRouter {
	r := new(ToDoRouter)
	r.handler = h
	r.router = mux.NewRouter()
	r.router.HandleFunc("/todo", r.handler.Create()).Methods(http.MethodPost)
	r.router.HandleFunc("/todo/{id}", r.handler.Read()).Methods(http.MethodGet)
	r.router.HandleFunc("/todo/{id}", r.handler.Update()).Methods(http.MethodPatch)
	r.router.HandleFunc("/todo/{id}", r.handler.Delete()).Methods(http.MethodDelete)
	r.router.HandleFunc("/todo", r.handler.List()).Methods(http.MethodGet)
	return r
}

func (r *ToDoRouter) GetRouter() *mux.Router {
	return r.router
}
