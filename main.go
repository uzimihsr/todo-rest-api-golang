package main

import (
	"fmt"
	"net/http"

	"github.com/uzimihsr/todo-rest-api-golang/presentation/handler"
	"github.com/uzimihsr/todo-rest-api-golang/presentation/router"
)

func main() {
	fmt.Println("Work in Progress")

	handler := handler.NewToDoHandler()
	router := router.NewToDoRouter(handler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.GetRouter(),
	}
	server.ListenAndServe()
}
