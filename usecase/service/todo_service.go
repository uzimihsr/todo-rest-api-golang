//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package service

import (
	"github.com/uzimihsr/todo-rest-api-golang/domain/model"
	"github.com/uzimihsr/todo-rest-api-golang/domain/repository"
)

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

func NewToDoService(repository repository.ToDoRepository) ToDoService {
	return &toDoService{repository: repository}
}

func (s *toDoService) Create(toDo *ToDoObject) (*ToDoObject, error) {

	createToDo := &model.ToDo{
		Title: toDo.Title,
		Done:  toDo.Done,
	}
	id, err := s.repository.Create(createToDo)
	if err != nil {
		return nil, err
	}

	result, err := s.repository.Read(id)
	if err != nil {
		return nil, err
	}

	return modelToObject(result), nil
}

func (s *toDoService) Read(toDo *ToDoObject) (*ToDoObject, error) {

	id := toDo.Id
	result, err := s.repository.Read(id)
	if err != nil {
		return nil, err
	}

	return modelToObject(result), nil
}

func (s *toDoService) Update(toDo *ToDoObject) (*ToDoObject, error) {

	before, err := s.repository.Read(toDo.Id)
	if err != nil {
		return nil, err
	}

	// 対象のToDoを更新
	updateToDo := &model.ToDo{
		Id:    toDo.Id,
		Title: toDo.Title,
		Done:  toDo.Done,
	}

	if updateToDo.Title == "" {
		updateToDo.Title = before.Title
	}
	err = s.repository.Update(updateToDo)
	if err != nil {
		return nil, err
	}

	// 更新されたToDoを取得
	result, err := s.repository.Read(toDo.Id)
	if err != nil {
		return nil, err
	}

	return modelToObject(result), nil
}

func (s *toDoService) Delete(toDo *ToDoObject) (*ToDoObject, error) {

	before, err := s.repository.Read(toDo.Id)
	if err != nil {
		return nil, err
	}

	// 対象のToDoを削除
	err = s.repository.Delete(toDo.Id)
	if err != nil {
		return nil, err
	}

	return modelToObject(before), nil
}

func (s *toDoService) List(option *ListOption) ([]ToDoObject, error) {

	selector := &model.ToDoSelector{}
	if option.Done != "" {
		selector.Done = option.Done
	} else {
		selector = nil
	}

	result, err := s.repository.List(selector)
	if err != nil {
		return nil, err
	}

	toDoList := []ToDoObject{}
	for _, t := range result {
		toDoList = append(toDoList, *modelToObject(&t))
	}

	return toDoList, nil
}

func modelToObject(model *model.ToDo) *ToDoObject {
	return &ToDoObject{
		Id:        model.Id,
		Title:     model.Title,
		Done:      model.Done,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
