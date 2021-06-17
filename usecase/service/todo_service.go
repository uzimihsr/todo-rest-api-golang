package service

import "github.com/uzimihsr/todo-rest-api-golang/domain/repository"

type ToDoService interface {
	Create(*ToDoObject) (*ToDoObject, error)
	Read(*ToDoObject) (*ToDoObject, error)
	Update(*ToDoObject) (*ToDoObject, error)
	Delete(*ToDoObject) (*ToDoObject, error)
	List(*ListOption) ([]ToDoObject, error)
}

type toDoService struct {
	repository repository.ToDoRepository
}

func NewToDoService() ToDoService {
	return new(toDoService)
}

func (s *toDoService) Create(toDo *ToDoObject) (*ToDoObject, error) {
	// WIP
	return nil, nil
}

func (s *toDoService) Read(toDo *ToDoObject) (*ToDoObject, error) {
	// WIP
	return nil, nil
}

func (s *toDoService) Update(toDo *ToDoObject) (*ToDoObject, error) {
	// WIP
	return nil, nil
}

func (s *toDoService) Delete(toDo *ToDoObject) (*ToDoObject, error) {
	// WIP
	return nil, nil
}

func (s *toDoService) List(option *ListOption) ([]ToDoObject, error) {
	// WIP
	return nil, nil
}
