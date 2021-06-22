//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package repository

import "github.com/uzimihsr/todo-rest-api-golang/domain/model"

type ToDoRepository interface {
	// Create new ToDo and return the ID
	Insert(*model.ToDo) (int64, error)

	// Read the ToDo specified by sthe ID
	SelectById(int64) (*model.ToDo, error)

	// Update the ToDo specified by the ID
	Update(*model.ToDo) error

	// Delete the ToDo specified by the ID
	DeleteById(int64) error

	// List all ToDo
	ListAll() ([]model.ToDo, error)

	// List by done status
	ListFilteredByDone(bool) ([]model.ToDo, error)
}
