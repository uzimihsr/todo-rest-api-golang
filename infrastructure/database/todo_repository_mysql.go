package database

import (
	"database/sql"
	"errors"
	"strconv"

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

func (r *toDoRepositoryMySQL) Create(model *model.ToDo) (int64, error) {
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

func (todoDB *toDoRepositoryMySQL) Read(id int64) (*model.ToDo, error) {
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

func (todoDB *toDoRepositoryMySQL) Delete(id int64) error {
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

func (todoDB *toDoRepositoryMySQL) List(selector *model.ToDoSelector) ([]model.ToDo, error) {
	var rows *sql.Rows
	var err error
	var toDoList []model.ToDo

	if selector == nil {
		rows, err = todoDB.db.Query("SELECT id, title, done, created_at, updated_at FROM todo")
		if err != nil {
			return nil, err
		}
	} else if selector.Done != "" {
		done, _ := strconv.ParseBool(selector.Done)
		rows, err = todoDB.db.Query("SELECT id, title, done, created_at, updated_at FROM todo WHERE done = ?", done)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid selector")
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
