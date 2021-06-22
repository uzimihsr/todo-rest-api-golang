package database

import (
	"database/sql"
	"errors"

	"github.com/uzimihsr/todo-rest-api-golang/domain/model"
	"github.com/uzimihsr/todo-rest-api-golang/domain/repository"
)

// implementation of repository
type toDoRepositoryMySQL struct {
	db *sql.DB
}

func NewToDoRepositoryMySQL(db *sql.DB) repository.ToDoRepository {
	return &toDoRepositoryMySQL{db: db}
}

func (r *toDoRepositoryMySQL) Insert(model *model.ToDo) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO todo(title, done) VALUES ( ?, ? )",
		model.Title,
		model.Done,
	)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (todoDB *toDoRepositoryMySQL) SelectById(id int64) (*model.ToDo, error) {
	todo := &model.ToDo{}
	err := todoDB.db.QueryRow(
		"SELECT id, title, done, created_at, updated_at FROM todo WHERE id = ?",
		id,
	).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Done,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (todoDB *toDoRepositoryMySQL) Update(model *model.ToDo) error {
	result, err := todoDB.db.Exec(
		"UPDATE todo SET title = ?, done = ? WHERE id = ?",
		model.Title,
		model.Done,
		model.Id,
	)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("UPDATE FAILED")
	}
	return nil
}

func (todoDB *toDoRepositoryMySQL) DeleteById(id int64) error {
	result, err := todoDB.db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("UPDATE FAILED")
	}
	return nil
}

func (todoDB *toDoRepositoryMySQL) ListAll() ([]model.ToDo, error) {
	var rows *sql.Rows
	var err error
	var toDoList []model.ToDo

	rows, err = todoDB.db.Query("SELECT id, title, done, created_at, updated_at FROM todo")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		todo := model.ToDo{}
		err := rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		toDoList = append(toDoList, todo)
	}

	return toDoList, nil
}

func (todoDB *toDoRepositoryMySQL) ListFilteredByDone(done bool) ([]model.ToDo, error) {
	var rows *sql.Rows
	var err error
	var toDoList []model.ToDo

	rows, err = todoDB.db.Query("SELECT id, title, done, created_at, updated_at FROM todo WHERE done = ?", done)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		todo := model.ToDo{}
		err := rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		toDoList = append(toDoList, todo)
	}

	return toDoList, nil
}
