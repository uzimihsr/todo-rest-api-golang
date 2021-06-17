//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package repository

import "github.com/uzimihsr/todo-rest-api-golang/domain/model"

type ToDoRepository interface {
	// Create new ToDo and return the ID of that
	Create(*model.ToDo) (int64, error)

	// Read the ToDo specified by ID
	Read(int64) (*model.ToDo, error)

	// Update the ToDo
	Update(*model.ToDo) error

	// Delete the ToDo specified by ID
	Delete(int64) error

	// List ToDo
	List(*model.ToDoSelector) ([]model.ToDo, error)
}
