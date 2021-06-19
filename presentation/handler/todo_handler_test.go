//go:generate mockgen -destination=../mock/mock_io.go -package=mock io Reader
package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/uzimihsr/todo-rest-api-golang/presentation/mock"
	"github.com/uzimihsr/todo-rest-api-golang/usecase/service"
	"github.com/uzimihsr/todo-rest-api-golang/usecase/service/mock_service"
)

func TestCreate(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	// Prepare
	ctrl := gomock.NewController(t)
	mockReader := mock.NewMockReader(ctrl)
	mockReader.EXPECT().Read(gomock.Any()).Return(0, errors.New("ERROR")).AnyTimes()
	requestJSON := &service.ToDoObject{
		Title: "test-ToDo",
	}
	j, _ := json.Marshal(requestJSON)
	tests := []struct {
		name               string
		createError        error
		createResult       *service.ToDoObject
		createTimes        int
		request            *http.Request
		expectedStatusCode int
	}{
		{
			name:               "01_正常にレスポンスが返せるケース",
			createError:        nil,
			createResult:       &service.ToDoObject{},
			createTimes:        1,
			expectedStatusCode: http.StatusOK,
			request:            httptest.NewRequest(http.MethodPost, "http://hogehoge/todo", bytes.NewBuffer(j)),
		},
		{
			name:               "02_parseRequestが失敗(ReadAllでエラー)するケース",
			createError:        nil,
			createResult:       nil,
			createTimes:        0,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodPost, "http://hogehoge/todo", mockReader),
		},
		{
			name:               "03_parseRequestが失敗(Unmarshalでエラー)するケース",
			createError:        nil,
			createResult:       nil,
			createTimes:        0,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodPost, "http://hogehoge/todo", bytes.NewBufferString("invalidRequestBody")),
		},
		{
			name:               "04_Createが失敗するケース",
			createError:        errors.New("Create ERROR"),
			createResult:       nil,
			createTimes:        1,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodPost, "http://hogehoge/todo", bytes.NewBuffer(j)),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			mockToDoService := mock_service.NewMockToDoService(ctrl)
			mockToDoService.EXPECT().Create(gomock.Any()).Return(tt.createResult, tt.createError).Times(tt.createTimes)
			toDoHandler := NewToDoHandler(mockToDoService)

			r := mux.NewRouter()
			r.HandleFunc("/todo", toDoHandler.Create()).Methods(http.MethodPost)
			w := httptest.NewRecorder()

			// Act
			r.ServeHTTP(w, tt.request)

			// Assert
			if w.Result().StatusCode != tt.expectedStatusCode {
				t.Errorf("expected: %d, actual: %d", tt.expectedStatusCode, w.Result().StatusCode)
			}
		})
	}
}

func TestRead(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	// Prepare
	ctrl := gomock.NewController(t)
	tests := []struct {
		name               string
		readError          error
		readResult         *service.ToDoObject
		readTimes          int
		request            *http.Request
		expectedStatusCode int
	}{
		{
			name:               "01_正常にレスポンスが返せるケース",
			readError:          nil,
			readResult:         &service.ToDoObject{},
			readTimes:          1,
			expectedStatusCode: http.StatusOK,
			request:            httptest.NewRequest(http.MethodGet, "http://hogehoge/todo/100", nil),
		},
		{
			name:               "02_getPathParamIdが失敗(Atoiでエラー)するケース",
			readError:          nil,
			readResult:         nil,
			readTimes:          0,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodGet, "http://hogehoge/todo/invalidId", nil),
		},
		{
			name:               "03_Readが失敗するケース",
			readError:          errors.New("Read ERROR"),
			readResult:         nil,
			readTimes:          1,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodGet, "http://hogehoge/todo/100", nil),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			mockToDoService := mock_service.NewMockToDoService(ctrl)
			mockToDoService.EXPECT().Read(gomock.Any()).Return(tt.readResult, tt.readError).Times(tt.readTimes)
			toDoHandler := NewToDoHandler(mockToDoService)

			r := mux.NewRouter()
			r.HandleFunc("/todo/{id}", toDoHandler.Read()).Methods(http.MethodGet)
			w := httptest.NewRecorder()

			// Act
			r.ServeHTTP(w, tt.request)

			// Assert
			if w.Result().StatusCode != tt.expectedStatusCode {
				t.Errorf("expected: %d, actual: %d", tt.expectedStatusCode, w.Result().StatusCode)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	// Prepare
	ctrl := gomock.NewController(t)
	mockReader := mock.NewMockReader(ctrl)
	mockReader.EXPECT().Read(gomock.Any()).Return(0, errors.New("ERROR")).AnyTimes()
	requestJSON := &service.ToDoObject{
		Id:    100,
		Title: "test-ToDo",
	}
	j, _ := json.Marshal(requestJSON)
	tests := []struct {
		name               string
		updateError        error
		updateResult       *service.ToDoObject
		updateTimes        int
		request            *http.Request
		expectedStatusCode int
	}{
		{
			name:               "01_正常にレスポンスが返せるケース",
			updateError:        nil,
			updateResult:       &service.ToDoObject{},
			updateTimes:        1,
			expectedStatusCode: http.StatusOK,
			request:            httptest.NewRequest(http.MethodPut, "http://hogehoge/todo/100", bytes.NewBuffer(j)),
		},
		{
			name:               "02_getPathParamIdが失敗(Atoiでエラー)するケース",
			updateError:        nil,
			updateResult:       nil,
			updateTimes:        0,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodPut, "http://hogehoge/todo/invalidId", bytes.NewBuffer(j)),
		},
		{
			name:               "03_parseRequestが失敗(ReadAllでエラー)するケース",
			updateError:        nil,
			updateResult:       nil,
			updateTimes:        0,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodPut, "http://hogehoge/todo/100", mockReader),
		},
		{
			name:               "04_parseRequestが失敗(Unmarshalでエラー)するケース",
			updateError:        nil,
			updateResult:       nil,
			updateTimes:        0,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodPut, "http://hogehoge/todo/100", bytes.NewBufferString("invalidRequestBody")),
		},
		{
			name:               "06_Updateが失敗するケース",
			updateError:        errors.New("Update ERROR"),
			updateResult:       nil,
			updateTimes:        1,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodPut, "http://hogehoge/todo/100", bytes.NewBuffer(j)),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			mockToDoService := mock_service.NewMockToDoService(ctrl)
			mockToDoService.EXPECT().Update(gomock.Any()).Return(tt.updateResult, tt.updateError).Times(tt.updateTimes)
			toDoHandler := NewToDoHandler(mockToDoService)

			r := mux.NewRouter()
			r.HandleFunc("/todo/{id}", toDoHandler.Update()).Methods(http.MethodPut)
			w := httptest.NewRecorder()

			// Act
			r.ServeHTTP(w, tt.request)

			// Assert
			if w.Result().StatusCode != tt.expectedStatusCode {
				t.Errorf("expected: %d, actual: %d", tt.expectedStatusCode, w.Result().StatusCode)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	// Prepare
	ctrl := gomock.NewController(t)
	tests := []struct {
		name               string
		deleteError        error
		deleteResult       *service.ToDoObject
		deleteTimes        int
		request            *http.Request
		expectedStatusCode int
	}{
		{
			name:               "01_正常にレスポンスが返せるケース",
			deleteError:        nil,
			deleteResult:       &service.ToDoObject{},
			deleteTimes:        1,
			expectedStatusCode: http.StatusOK,
			request:            httptest.NewRequest(http.MethodDelete, "http://hogehoge/todo/100", nil),
		},
		{
			name:               "02_getPathParamIdが失敗(Atoiでエラー)するケース",
			deleteError:        nil,
			deleteResult:       nil,
			deleteTimes:        0,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodDelete, "http://hogehoge/todo/invalidId", nil),
		},
		{
			name:               "03_Deleteが失敗するケース",
			deleteError:        errors.New("Read ERROR"),
			deleteResult:       nil,
			deleteTimes:        1,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodDelete, "http://hogehoge/todo/100", nil),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			mockToDoService := mock_service.NewMockToDoService(ctrl)
			mockToDoService.EXPECT().Delete(gomock.Any()).Return(tt.deleteResult, tt.deleteError).Times(tt.deleteTimes)
			toDoHandler := NewToDoHandler(mockToDoService)

			r := mux.NewRouter()
			r.HandleFunc("/todo/{id}", toDoHandler.Delete()).Methods(http.MethodDelete)
			w := httptest.NewRecorder()

			// Act
			r.ServeHTTP(w, tt.request)

			// Assert
			if w.Result().StatusCode != tt.expectedStatusCode {
				t.Errorf("expected: %d, actual: %d", tt.expectedStatusCode, w.Result().StatusCode)
			}
		})
	}
}

func TestList(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	// Prepare
	ctrl := gomock.NewController(t)
	tests := []struct {
		name               string
		listError          error
		listResult         []service.ToDoObject
		listTimes          int
		request            *http.Request
		expectedStatusCode int
	}{
		{
			name:               "01_正常にレスポンスが返せるケース",
			listError:          nil,
			listResult:         []service.ToDoObject{{Id: 100}},
			listTimes:          1,
			expectedStatusCode: http.StatusOK,
			request:            httptest.NewRequest(http.MethodGet, "http://hogehoge/todo", nil),
		},
		{
			name:               "03_Listが失敗するケース",
			listError:          errors.New("Read ERROR"),
			listResult:         nil,
			listTimes:          1,
			expectedStatusCode: http.StatusInternalServerError,
			request:            httptest.NewRequest(http.MethodGet, "http://hogehoge/todo", nil),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			mockToDoService := mock_service.NewMockToDoService(ctrl)
			mockToDoService.EXPECT().List(gomock.Any()).Return(tt.listResult, tt.listError).Times(tt.listTimes)
			toDoHandler := NewToDoHandler(mockToDoService)

			r := mux.NewRouter()
			r.HandleFunc("/todo", toDoHandler.List()).Methods(http.MethodGet)
			w := httptest.NewRecorder()

			// Act
			r.ServeHTTP(w, tt.request)

			// Assert
			if w.Result().StatusCode != tt.expectedStatusCode {
				t.Errorf("expected: %d, actual: %d", tt.expectedStatusCode, w.Result().StatusCode)
			}
		})
	}
}
