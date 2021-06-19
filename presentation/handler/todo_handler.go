package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func NewToDoHandler(service service.ToDoService) ToDoHandler {
	return &toDoHandler{
		service: service,
	}
}

func (h *toDoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestToDo, err := parseRequestJSON(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resultToDo, err := h.service.Create(requestToDo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultToDo)
	}
}

func (h *toDoHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getPathParamId(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requestToDo := &service.ToDoObject{
			Id: id,
		}

		resultToDo, err := h.service.Read(requestToDo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultToDo)
	}
}

func (h *toDoHandler) Update() http.HandlerFunc {
	// WIP
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getPathParamId(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		requestToDo, err := parseRequestJSON(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requestToDo.Id = id

		// DBのレコードを更新
		resultToDo, err := h.service.Update(requestToDo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultToDo)
	}
}

func (h *toDoHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getPathParamId(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requestToDo := &service.ToDoObject{
			Id: id,
		}

		// DBのレコードを削除
		resultToDo, err := h.service.Delete(requestToDo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultToDo)
	}
}

func (h *toDoHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		done := r.FormValue("done")
		listOption := &service.ListOption{
			Done: done,
		}
		todoList, err := h.service.List(listOption)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resultList := []service.ToDoObject{}
		resultList = append(resultList, todoList...)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultList)
	}
}

// パスパラメータ{id}を取得する
func getPathParamId(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return -1, err
	}
	return int64(id), nil
}

// リクエストボディをJSONにパースする
func parseRequestJSON(r *http.Request) (*service.ToDoObject, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	reqObject := &service.ToDoObject{}
	err = json.Unmarshal(reqBody, reqObject)
	if err != nil {
		return nil, err
	}
	return reqObject, nil
}
