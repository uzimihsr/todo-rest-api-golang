package service

import (
	"errors"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/uzimihsr/todo-rest-api-golang/domain/model"
	"github.com/uzimihsr/todo-rest-api-golang/domain/repository/mock_repository"
)

func TestCreate(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	toDoObject := &ToDoObject{
		Title: "test-ToDo",
		Done:  false,
	}

	tests := []struct {
		name        string
		createError error
		createId    int64
		createTimes int
		readError   error
		readResult  *model.ToDo
		readTimes   int
		wantError   bool
	}{
		{
			name:        "01_CreateとReadが成功するケース",
			createError: nil,
			createId:    100,
			createTimes: 1,
			readError:   nil,
			readResult:  &model.ToDo{Title: toDoObject.Title, Done: toDoObject.Done},
			readTimes:   1,
			wantError:   false,
		},
		{
			name:        "02_Createが失敗するケース",
			createError: errors.New("Create ERROR"),
			createId:    -1,
			createTimes: 1,
			readError:   nil,
			readResult:  nil,
			readTimes:   0,
			wantError:   true,
		},
		{
			name:        "03_Readが失敗するケース",
			createError: nil,
			createId:    100,
			createTimes: 1,
			readError:   errors.New("Read ERROR"),
			readResult:  nil,
			readTimes:   1,
			wantError:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			ctrl := gomock.NewController(t)
			mockToDoRepository := mock_repository.NewMockToDoRepository(ctrl)
			mockToDoRepository.EXPECT().Insert(gomock.Any()).Return(tt.createId, tt.createError).Times(tt.createTimes)
			mockToDoRepository.EXPECT().SelectById(gomock.Any()).Return(tt.readResult, tt.readError).Times(tt.readTimes)
			toDoService := NewToDoService(mockToDoRepository)

			// Act
			result, err := toDoService.Create(toDoObject)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
			if (result != nil) && (result.Title != toDoObject.Title || result.Done != toDoObject.Done) {
				t.Errorf("values do not match.\n expected(title): %v, actual(title): %v \n expected(done): %v, actual(done): %v", toDoObject.Title, result.Title, toDoObject.Done, result.Done)
			}
		})
	}
}

func TestRead(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	toDoObject := &ToDoObject{
		Id: 100,
	}
	tests := []struct {
		name       string
		readError  error
		readResult *model.ToDo
		readTimes  int
		wantError  bool
	}{
		{
			name:       "01_Readが成功するケース",
			readError:  nil,
			readResult: &model.ToDo{Id: toDoObject.Id},
			readTimes:  1,
			wantError:  false,
		},
		{
			name:       "02_Readが失敗するケース",
			readError:  errors.New("Read ERROR"),
			readResult: nil,
			readTimes:  1,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			ctrl := gomock.NewController(t)
			mockToDoRepository := mock_repository.NewMockToDoRepository(ctrl)
			mockToDoRepository.EXPECT().SelectById(gomock.Any()).Return(tt.readResult, tt.readError).Times(tt.readTimes)
			toDoService := NewToDoService(mockToDoRepository)

			// Act
			result, err := toDoService.Read(toDoObject)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
			if (result != nil) && (result.Id != toDoObject.Id) {
				t.Errorf("value does not match.\n expected(id): %v, actual(id): %v", toDoObject.Id, result.Id)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	toDoObject := &ToDoObject{
		Id:    100,
		Title: "test-ToDo",
		Done:  true,
	}

	tests := []struct {
		name        string
		readError1  error
		readResult1 *model.ToDo
		readTimes1  int
		updateError error
		updateTimes int
		readError2  error
		readResult2 *model.ToDo
		readTimes2  int
		wantError   bool
	}{
		{
			name:        "01_Read(1回目)+Update+Read(2回目)が成功するケース",
			readError1:  nil,
			readResult1: &model.ToDo{},
			readTimes1:  1,
			updateError: nil,
			updateTimes: 1,
			readError2:  nil,
			readResult2: &model.ToDo{Id: toDoObject.Id, Title: toDoObject.Title, Done: toDoObject.Done},
			readTimes2:  1,
			wantError:   false,
		},
		{
			name:        "02_Read(1回目)が失敗するケース",
			readError1:  errors.New("Read ERROR"),
			readResult1: nil,
			readTimes1:  1,
			updateError: nil,
			updateTimes: 0,
			readError2:  nil,
			readResult2: nil,
			readTimes2:  0,
			wantError:   true,
		},
		{
			name:        "03_Updateが失敗するケース",
			readError1:  nil,
			readResult1: &model.ToDo{},
			readTimes1:  1,
			updateError: errors.New("Update ERROR"),
			updateTimes: 1,
			readError2:  nil,
			readResult2: nil,
			readTimes2:  0,
			wantError:   true,
		},
		{
			name:        "04_Read(2回目)が失敗するケース",
			readError1:  nil,
			readResult1: &model.ToDo{},
			readTimes1:  1,
			updateError: nil,
			updateTimes: 1,
			readError2:  errors.New("Read ERROR"),
			readResult2: nil,
			readTimes2:  1,
			wantError:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			ctrl := gomock.NewController(t)
			mockToDoRepository := mock_repository.NewMockToDoRepository(ctrl)
			gomock.InOrder(
				mockToDoRepository.EXPECT().SelectById(gomock.Any()).Return(tt.readResult1, tt.readError1).Times(tt.readTimes1),
				mockToDoRepository.EXPECT().SelectById(gomock.Any()).Return(tt.readResult2, tt.readError2).Times(tt.readTimes2),
			)
			mockToDoRepository.EXPECT().Update(gomock.Any()).Return(tt.updateError).Times(tt.updateTimes)
			toDoService := NewToDoService(mockToDoRepository)

			// Act
			result, err := toDoService.Update(toDoObject)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
			if (result != nil) && (result.Title != toDoObject.Title || result.Done != toDoObject.Done) {
				t.Errorf("values do not match.\n expected(title): %v, actual(title): %v \n expected(done): %v, actual(done): %v", toDoObject.Title, result.Title, toDoObject.Done, result.Done)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	toDoObject := &ToDoObject{
		Id: 100,
	}

	tests := []struct {
		name        string
		readError   error
		readResult  *model.ToDo
		readTimes   int
		deleteError error
		deleteTimes int
		wantError   bool
	}{
		{
			name:        "01_Read+Deleteが成功するケース",
			readError:   nil,
			readResult:  &model.ToDo{Id: toDoObject.Id},
			readTimes:   1,
			deleteError: nil,
			deleteTimes: 1,
			wantError:   false,
		},
		{
			name:        "02_Readが失敗するケース",
			readError:   errors.New("Read ERROR"),
			readResult:  nil,
			readTimes:   1,
			deleteError: nil,
			deleteTimes: 0,
			wantError:   true,
		},
		{
			name:        "02_Deleteが失敗するケース",
			readError:   nil,
			readResult:  &model.ToDo{},
			readTimes:   1,
			deleteError: errors.New("Delete ERROR"),
			deleteTimes: 1,
			wantError:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			ctrl := gomock.NewController(t)
			mockToDoRepository := mock_repository.NewMockToDoRepository(ctrl)
			mockToDoRepository.EXPECT().SelectById(gomock.Any()).Return(tt.readResult, tt.readError).Times(tt.readTimes)
			mockToDoRepository.EXPECT().DeleteById(gomock.Any()).Return(tt.deleteError).Times(tt.deleteTimes)
			toDoService := NewToDoService(mockToDoRepository)

			// Act
			result, err := toDoService.Delete(toDoObject)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
			if (result != nil) && (result.Id != toDoObject.Id) {
				t.Errorf("value does not match.\n expected(id): %v, actual(id): %v", toDoObject.Id, result.Id)
			}
		})
	}
}

func TestList(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	toDoList := []model.ToDo{
		{
			Id:    100,
			Title: "test-ToDo01",
			Done:  false,
		},
		{
			Id:    101,
			Title: "test-ToDo02",
			Done:  true,
		},
	}

	tests := []struct {
		name                    string
		listOption              ListOption
		listError               error
		listResult              []model.ToDo
		listAllTimes            int
		listFilteredByDoneTimes int
		wantError               bool
	}{
		{
			name:                    "01_Listが成功するケース_Doneが指定されている場合",
			listOption:              ListOption{Done: "true"},
			listError:               nil,
			listResult:              toDoList,
			listAllTimes:            0,
			listFilteredByDoneTimes: 1,
			wantError:               false,
		},
		{
			name:                    "02_Listが成功するケース_doneが指定されていない場合",
			listOption:              ListOption{Done: ""},
			listError:               nil,
			listResult:              toDoList,
			listAllTimes:            1,
			listFilteredByDoneTimes: 0,
			wantError:               false,
		},
		{
			name:                    "03_Listが失敗するケース_doneが指定されている場合",
			listOption:              ListOption{Done: "false"},
			listError:               errors.New("List ERROR"),
			listResult:              nil,
			listAllTimes:            0,
			listFilteredByDoneTimes: 1,
			wantError:               true,
		},
		{
			name:                    "04_Listが失敗するケース_doneが指定されていない場合",
			listOption:              ListOption{Done: ""},
			listError:               errors.New("List ERROR"),
			listResult:              nil,
			listAllTimes:            1,
			listFilteredByDoneTimes: 0,
			wantError:               true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			var done bool
			if tt.listOption.Done != "" {
				done, _ = strconv.ParseBool(tt.listOption.Done)
			}
			ctrl := gomock.NewController(t)
			mockToDoRepository := mock_repository.NewMockToDoRepository(ctrl)
			mockToDoRepository.EXPECT().ListAll().Return(tt.listResult, tt.listError).Times(tt.listAllTimes)
			mockToDoRepository.EXPECT().ListFilteredByDone(done).Return(tt.listResult, tt.listError).Times(tt.listFilteredByDoneTimes)
			toDoService := NewToDoService(mockToDoRepository)

			// Act
			result, err := toDoService.List(&tt.listOption)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
			if (result != nil) && (len(result) != len(toDoList)) {
				t.Errorf("lengths do not match. expected: %v, actual: %v", len(toDoList), len(result))
			}
		})
	}
}
