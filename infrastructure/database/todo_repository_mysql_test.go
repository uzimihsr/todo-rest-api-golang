package database

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/uzimihsr/todo-rest-api-golang/domain/model"
)

// with sqlmock
func TestInsert(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	tests := []struct {
		name       string
		execError  error
		execResult driver.Result
		wantError  bool
	}{
		{
			name:       "01_INSERTが成功するケース",
			execError:  nil,
			execResult: sqlmock.NewResult(1, 1),
			wantError:  false,
		},
		{
			name:       "02_Execでエラーが発生して失敗するケース",
			execError:  errors.New("INSERT FAILED"),
			execResult: sqlmock.NewErrorResult(errors.New("ERROR RESULT")),
			wantError:  true,
		},
		{
			name:       "03_LastInsertIdが取得できず失敗するケース",
			execError:  nil,
			execResult: sqlmock.NewErrorResult(errors.New("ERROR RESULT")),
			wantError:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			toDoModel := &model.ToDo{
				Title: "test-ToDo",
				Done:  false,
			}
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO todo(title, done) VALUES ( ?, ? )")).
				WithArgs(toDoModel.Title, toDoModel.Done).
				WillReturnResult(tt.execResult).
				WillReturnError(tt.execError)
			toDoRepository := NewToDoRepositoryMySQL(db)

			// Act
			_, err = toDoRepository.Insert(toDoModel)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
		})
	}
}

func TestSelectById(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	tests := []struct {
		name      string
		queryRow  *sqlmock.Rows
		wantError bool
	}{
		{
			name:      "01_SELECTが成功するケース",
			queryRow:  sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}).AddRow(1, "test-ToDo", false, time.Now(), time.Now()),
			wantError: false,
		},
		{
			name:      "02_Scanが失敗するケース",
			queryRow:  sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}),
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			id := int64(100)
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, done, created_at, updated_at FROM todo WHERE id = ?")).
				WithArgs(id).
				WillReturnRows(tt.queryRow)
			toDoRepository := NewToDoRepositoryMySQL(db)

			// Act
			_, err = toDoRepository.SelectById(id)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	tests := []struct {
		name       string
		execError  error
		execResult driver.Result
		wantError  bool
	}{
		{
			name:       "01_UPDATEが成功するケース",
			execError:  nil,
			execResult: sqlmock.NewResult(1, 1),
			wantError:  false,
		},
		{
			name:       "02_UPDATEが失敗するケース",
			execError:  errors.New("UPDATE FAILED"),
			execResult: sqlmock.NewErrorResult(errors.New("ERROR RESULT")),
			wantError:  true,
		},
		{
			name:       "03_RowsAffected()が失敗するケース",
			execError:  nil,
			execResult: sqlmock.NewErrorResult(errors.New("ERROR RESULT")),
			wantError:  true,
		},
		{
			name:       "04_RowsAffected()で1以外が返るケース",
			execError:  nil,
			execResult: sqlmock.NewResult(0, 0),
			wantError:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			toDoModel := &model.ToDo{
				Id:    100,
				Title: "test-ToDo",
				Done:  true,
			}
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("UPDATE todo SET title = ?, done = ? WHERE id = ?")).
				WithArgs(toDoModel.Title, toDoModel.Done, toDoModel.Id).
				WillReturnResult(tt.execResult).
				WillReturnError(tt.execError)
			toDoRepository := NewToDoRepositoryMySQL(db)

			// Act
			err = toDoRepository.Update(toDoModel)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
		})
	}
}

func TestDeleteById(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	tests := []struct {
		name       string
		execError  error
		execResult driver.Result
		wantError  bool
	}{
		{
			name:       "01_DELETEが成功するケース",
			execError:  nil,
			execResult: sqlmock.NewResult(0, 1),
			wantError:  false,
		},
		{
			name:       "02_DELETEが失敗するケース",
			execError:  errors.New("DELETE FAILED"),
			execResult: nil,
			wantError:  true,
		},
		{
			name:       "03_RowsAffected()が失敗するケース",
			execError:  nil,
			execResult: sqlmock.NewErrorResult(errors.New("ERROR RESULT")),
			wantError:  true,
		},
		{
			name:       "04_RowsAffected()が1以外を返すケース",
			execError:  nil,
			execResult: sqlmock.NewResult(0, 100),
			wantError:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			id := int64(100)
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo WHERE id = ?")).
				WithArgs(id).
				WillReturnResult(tt.execResult).
				WillReturnError(tt.execError)
			toDoRepository := NewToDoRepositoryMySQL(db)

			// Act
			err = toDoRepository.DeleteById(id)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
		})
	}
}

func TestListAll(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	tests := []struct {
		name       string
		queryRow   *sqlmock.Rows
		queryError error
		wantError  bool
	}{
		{
			name:       "01_SELECTが成功するケース",
			queryRow:   sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}).AddRow(1, "test-ToDo", true, time.Now(), time.Now()),
			queryError: nil,
			wantError:  false,
		},
		{
			name:       "02_SELECTが失敗するケース",
			queryRow:   sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}),
			queryError: errors.New("SELECT FAILED"),
			wantError:  true,
		},
		{
			name:       "03_Scanが失敗するケース",
			queryRow:   sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}).AddRow(nil, nil, nil, nil, nil),
			queryError: nil,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, done, created_at, updated_at FROM todo")).
				WillReturnRows(tt.queryRow).
				WillReturnError(tt.queryError)
			toDoRepository := NewToDoRepositoryMySQL(db)

			// Act
			_, err = toDoRepository.ListAll()

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
		})
	}
}

func TestListFilteredByDone(t *testing.T) {
	t.Parallel() // https://github.com/golang/go/wiki/TableDrivenTests

	tests := []struct {
		name       string
		done       bool
		queryRow   *sqlmock.Rows
		queryError error
		wantError  bool
	}{
		{
			name:       "01_SELECTが成功するケース_done=true",
			done:       true,
			queryRow:   sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}).AddRow(1, "test-ToDo", true, time.Now(), time.Now()),
			queryError: nil,
			wantError:  false,
		},
		{
			name:       "02_SELECTが成功するケース_done=false",
			done:       false,
			queryRow:   sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}).AddRow(1, "test-ToDo", false, time.Now(), time.Now()),
			queryError: nil,
			wantError:  false,
		},
		{
			name:       "03_SELECTが失敗するケース",
			done:       true,
			queryRow:   sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}),
			queryError: errors.New("SELECT FAILED"),
			wantError:  true,
		},
		{
			name:       "04_Scanが失敗するケース",
			done:       true,
			queryRow:   sqlmock.NewRows([]string{"id", "title", "done", "created_at", "updated_at"}).AddRow(nil, nil, nil, nil, nil),
			queryError: nil,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)

			// Arrange
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, done, created_at, updated_at FROM todo WHERE done = ?")).
				WithArgs(tt.done).
				WillReturnRows(tt.queryRow).
				WillReturnError(tt.queryError)
			toDoRepository := NewToDoRepositoryMySQL(db)

			// Act
			_, err = toDoRepository.ListFilteredByDone(tt.done)

			// Assert
			if (err != nil) != tt.wantError {
				t.Error(err.Error())
			}
		})
	}
}

func TestInsertWithDB(t *testing.T) {
	t.Parallel()

	// Arrange
	resource, pool := createMySQLContainer("test/table_without_records.sql")
	defer closeMySQLContainer(resource, pool)
	db := connectMySQLContainer(resource, pool)
	expected := &model.ToDo{
		Title: "testToDo",
	}
	toDoRepository := NewToDoRepositoryMySQL(db)

	// Act
	id, err := toDoRepository.Insert(expected)
	if err != nil {
		t.Error(err.Error())
	}

	// Assert
	actual := &model.ToDo{}
	err = db.QueryRow("SELECT id, title, done FROM todo WHERE id = ?", id).Scan(&actual.Id, &actual.Title, &actual.Done)
	if err != nil {
		t.Error(err.Error())
	}
	if actual.Title != expected.Title {
		t.Errorf("expected: %s, actual: %s", expected.Title, actual.Title)
	}
	if actual.Done != expected.Done {
		t.Errorf("expected: %v, actual: %v", expected.Done, actual.Done)
	}
}

func TestSelectByIdWithDB(t *testing.T) {
	t.Parallel()

	// Arrange
	resource, pool := createMySQLContainer("test/table_with_records.sql")
	defer closeMySQLContainer(resource, pool)
	db := connectMySQLContainer(resource, pool)
	expected := &model.ToDo{
		Title: "ToDo01", // see test/test_read.sql
		Done:  false,
	}
	toDoRepository := NewToDoRepositoryMySQL(db)

	// Act
	actual, err := toDoRepository.SelectById(1)
	if err != nil {
		t.Error(err.Error())
	}

	// Assert
	if actual.Title != expected.Title {
		t.Errorf("expected: %s, actual: %s", expected.Title, actual.Title)
	}
	if actual.Done != expected.Done {
		t.Errorf("expected: %v, actual: %v", expected.Done, actual.Done)
	}

}

func TestUpdateWithDB(t *testing.T) {
	t.Parallel()
	// Arrange
	resource, pool := createMySQLContainer("test/table_with_records.sql")
	defer closeMySQLContainer(resource, pool)
	db := connectMySQLContainer(resource, pool)
	expected := &model.ToDo{
		Id:    1, // see test/test_read.sql
		Title: "ToDo01",
		Done:  true,
	}
	toDoRepository := NewToDoRepositoryMySQL(db)

	// Act
	err := toDoRepository.Update(expected)
	if err != nil {
		t.Error(err.Error())
	}

	// Assert
	actual := &model.ToDo{}
	err = db.QueryRow("SELECT id, title, done FROM todo WHERE id = ?", expected.Id).Scan(&actual.Id, &actual.Title, &actual.Done)
	if err != nil {
		t.Error(err.Error())
	}
	if actual.Title != expected.Title {
		t.Errorf("expected: %s, actual: %s", expected.Title, actual.Title)
	}
	if actual.Done != expected.Done {
		t.Errorf("expected: %v, actual: %v", expected.Done, actual.Done)
	}

}

func TestDeleteByIdWithDB(t *testing.T) {
	t.Parallel()
	// Arrange
	resource, pool := createMySQLContainer("test/table_with_records.sql")
	defer closeMySQLContainer(resource, pool)
	db := connectMySQLContainer(resource, pool)
	toDoRepository := NewToDoRepositoryMySQL(db)

	// Act
	err := toDoRepository.DeleteById(1)
	if err != nil {
		t.Error(err.Error())
	}

	// Assert
	actual := &model.ToDo{}
	err = db.QueryRow("SELECT id, title, done FROM todo WHERE id = ?", 1).Scan(&actual.Id, &actual.Title, &actual.Done)
	if err == nil {
		t.Error("THE RECORD STILL EXISTS")
	}

}

func TestListAllWithDB(t *testing.T) {
	t.Parallel()

	// Arrange
	resource, pool := createMySQLContainer("test/table_with_records.sql")
	defer closeMySQLContainer(resource, pool)
	db := connectMySQLContainer(resource, pool)
	toDoRepository := NewToDoRepositoryMySQL(db)
	expected := []model.ToDo{ // see test/test_read.sql
		{
			Id:    1,
			Title: "ToDo01",
			Done:  false,
		},
		{
			Id:    2,
			Title: "ToDo02",
			Done:  false,
		},
		{
			Id:    3,
			Title: "ToDo03",
			Done:  true,
		},
		{
			Id:    4,
			Title: "ToDo04",
			Done:  true,
		},
		{
			Id:    5,
			Title: "ToDo05",
			Done:  false,
		},
	}

	// Act
	actual, err := toDoRepository.ListAll()
	if err != nil {
		t.Error(err.Error())
	}

	// Assert
	if len(actual) != len(expected) {
		t.Errorf("list lengths do not match. expected: %v, actual: %v", len(expected), len(actual))
	}
	for i := range actual {
		err := checkRecord(&expected[i], &actual[i])
		if err != nil {
			t.Error(err.Error())
		}
	}
}

// with dockertest
func TestListFilteredByDoneWithDB(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			"ListFilteredById(done=true)でMySQLからSELECTしたレコードのチェック",
			func(t *testing.T) {
				t.Parallel()
				// Arrange
				resource, pool := createMySQLContainer("test/table_with_records.sql")
				defer closeMySQLContainer(resource, pool)
				db := connectMySQLContainer(resource, pool)
				toDoRepository := NewToDoRepositoryMySQL(db)
				expected := []model.ToDo{ // see test/test_read.sql
					{
						Id:    3,
						Title: "ToDo03",
						Done:  true,
					},
					{
						Id:    4,
						Title: "ToDo04",
						Done:  true,
					},
				}

				// Act
				actual, err := toDoRepository.ListFilteredByDone(true)
				if err != nil {
					t.Error(err.Error())
				}

				// Assert
				if len(actual) != len(expected) {
					t.Errorf("list lengths do not match. expected: %v, actual: %v", len(expected), len(actual))
				}
				for i := range actual {
					err := checkRecord(&expected[i], &actual[i])
					if err != nil {
						t.Error(err.Error())
					}
				}
			},
		},
		{
			"ListFilteredById(done=false)でMySQLからSELECTしたレコードのチェック",
			func(t *testing.T) {
				t.Parallel()
				// Arrange
				resource, pool := createMySQLContainer("test/table_with_records.sql")
				defer closeMySQLContainer(resource, pool)
				db := connectMySQLContainer(resource, pool)
				toDoRepository := NewToDoRepositoryMySQL(db)
				expected := []model.ToDo{ // see test/test_read.sql
					{
						Id:    1,
						Title: "ToDo01",
						Done:  false,
					},
					{
						Id:    2,
						Title: "ToDo02",
						Done:  false,
					},
					{
						Id:    5,
						Title: "ToDo05",
						Done:  false,
					},
				}

				// Act
				actual, err := toDoRepository.ListFilteredByDone(false)
				if err != nil {
					t.Error(err.Error())
				}

				// Assert
				if len(actual) != len(expected) {
					t.Errorf("list lengths do not match. expected: %v, actual: %v", len(expected), len(actual))
				}
				for i := range actual {
					err := checkRecord(&expected[i], &actual[i])
					if err != nil {
						t.Error(err.Error())
					}
				}
			},
		},
	}
	for _, tt := range tests {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.name, tt.fn)
	}

	// t.Run(
	// 	"ListFilteredById(done=true)でMySQLからSELECTしたレコードのチェック",
	// 	func(t *testing.T) {
	// 		// Arrange
	// 		resource, pool := createMySQLContainer("test/table_with_records.sql")
	// 		defer closeMySQLContainer(resource, pool)
	// 		db := connectMySQLContainer(resource, pool)
	// 		toDoRepository := NewToDoRepositoryMySQL(db)
	// 		expected := []model.ToDo{ // see test/test_read.sql
	// 			{
	// 				Id:    3,
	// 				Title: "ToDo03",
	// 				Done:  true,
	// 			},
	// 			{
	// 				Id:    4,
	// 				Title: "ToDo04",
	// 				Done:  true,
	// 			},
	// 		}

	// 		// Act
	// 		actual, err := toDoRepository.ListFilteredByDone(true)
	// 		if err != nil {
	// 			t.Error(err.Error())
	// 		}

	// 		// Assert
	// 		if len(actual) != len(expected) {
	// 			t.Errorf("list lengths do not match. expected: %v, actual: %v", len(expected), len(actual))
	// 		}
	// 		for i := range actual {
	// 			err := checkRecord(&expected[i], &actual[i])
	// 			if err != nil {
	// 				t.Error(err.Error())
	// 			}
	// 		}
	// 	},
	// )

	// t.Run(
	// 	"ListFilteredById(done=false)でMySQLからSELECTしたレコードのチェック",
	// 	func(t *testing.T) {
	// 		// Arrange
	// 		resource, pool := createMySQLContainer("test/table_with_records.sql")
	// 		defer closeMySQLContainer(resource, pool)
	// 		db := connectMySQLContainer(resource, pool)
	// 		toDoRepository := NewToDoRepositoryMySQL(db)
	// 		expected := []model.ToDo{ // see test/test_read.sql
	// 			{
	// 				Id:    1,
	// 				Title: "ToDo01",
	// 				Done:  false,
	// 			},
	// 			{
	// 				Id:    2,
	// 				Title: "ToDo02",
	// 				Done:  false,
	// 			},
	// 			{
	// 				Id:    5,
	// 				Title: "ToDo05",
	// 				Done:  false,
	// 			},
	// 		}

	// 		// Act
	// 		actual, err := toDoRepository.ListFilteredByDone(false)
	// 		if err != nil {
	// 			t.Error(err.Error())
	// 		}

	// 		// Assert
	// 		if len(actual) != len(expected) {
	// 			t.Errorf("list lengths do not match. expected: %v, actual: %v", len(expected), len(actual))
	// 		}
	// 		for i := range actual {
	// 			err := checkRecord(&expected[i], &actual[i])
	// 			if err != nil {
	// 				t.Error(err.Error())
	// 			}
	// 		}
	// 	},
	// )

}

// Create Docker container for tests
func createMySQLContainer(sqlFileName string) (*dockertest.Resource, *dockertest.Pool) {
	// Dockerコンテナへのファイルマウント時に絶対パスが必要
	pwd, _ := os.Getwd()

	// connect to docker
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// mysql options
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7.33",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=secret",
		},
		Mounts: []string{
			pwd + "/test/my.cnf:/etc/mysql/my.cnf",
			pwd + "/" + sqlFileName + ":/docker-entrypoint-initdb.d/schema.sql", //コンテナ起動時に流し込むスクリプト
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	// start container
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func closeMySQLContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// stop container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

// connect to the container
func connectMySQLContainer(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {

	var db *sql.DB
	if err := pool.Retry(func() error {
		// wait for container setup
		time.Sleep(time.Second * 10)

		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/todo_db?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}

func checkRecord(expected *model.ToDo, actual *model.ToDo) error {
	if actual.Title != expected.Title {
		return fmt.Errorf("expected: %s, actual: %s", expected.Title, actual.Title)
	}
	if actual.Done != expected.Done {
		return fmt.Errorf("expected: %v, actual: %v", expected.Done, actual.Done)
	}
	return nil
}
