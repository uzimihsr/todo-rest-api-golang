package database

import (
	"github.com/uzimihsr/todo-rest-api-golang/domain/model"
	"github.com/uzimihsr/todo-rest-api-golang/domain/repository"
)

// implementation of repository
type toDoRepositoryMySQL struct {
}

func NewToDoRepositoryMySQL() repository.ToDoRepository {
	t := new(toDoRepositoryMySQL)
	return t
}

func (todoDB *toDoRepositoryMySQL) Create(model *model.ToDo) (int64, error) {
	// WIP
	return 0, nil
}

func (todoDB *toDoRepositoryMySQL) Read(id int64) (*model.ToDo, error) {
	// WIP
	return nil, nil
}

func (todoDB *toDoRepositoryMySQL) Update(model *model.ToDo) error {
	// WIP
	return nil
}

func (todoDB *toDoRepositoryMySQL) Delete(id int64) error {
	// WIP
	return nil
}

func (todoDB *toDoRepositoryMySQL) List(selector *model.ToDoSelector) ([]model.ToDo, error) {
	// WIP
	return nil, nil
}
