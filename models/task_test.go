package models

import (
	"fmt"
	"go-gorm-gin/config"
	"go-gorm-gin/models"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var deadline, _ = time.Parse(time.RFC3339, "2020-06-18T13:13:23.63608771-04:00")

type Tests struct {
	ID          int
	Name        string
	Description string
	Deadline    time.Time
	Taskstatus  string
}

var testTask = []Tests{
	{
		ID:          1,
		Name:        "Register",
		Description: "Register for conference",
		Deadline:    deadline,
		Taskstatus:  "Active",
	},
}

func initmockdb() (mock sqlmock.Sqlmock, rows *sqlmock.Rows, err error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()
	rows = sqlmock.NewRows([]string{"id", "name", "description", "deadline", "taskstatus"}).
		AddRow(1, "Register", "Register for conference", "2020-06-18T13:13:23.63608771-04:00", "Active")
	//AddRow(1, "Renew", "Renew the license", "2020-07-8T13:13:23.63608771-03:00", "Active")
	return
}
func TestGetAllTasks(t *testing.T) {

	mock, rows, err := initmockdb()
	mock.ExpectQuery(`SELECT * FROM "tasks"`).WillReturnRows(rows)

	var testTask *[]models.Task
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Fatal("Status:", err)
	}
	defer config.DB.Close()

	if err = models.GetAllTasks(testTask); err != nil {
		t.Logf(err.Error())
	}
	if err = config.DB.Find(&testTask).Error; err != nil {
		t.Logf(err.Error())
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expections: %s", err)
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestCreateTask(t *testing.T) {
	mock, rows, err := initmockdb()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks"`)).WithArgs(rows)
	//var testTask1 *[]models.Task
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Fatal("Status:", err)
	}
	defer config.DB.Close()
	//testTask1 = testTask
	// if err = models.CreateTask(testTask1); err != nil {
	// 	t.Logf(err.Error())
	// }
	if err = config.DB.Create(&testTask).Error; err != nil {
		t.Logf(err.Error())
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expections: %s", err)
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestGetTaskByID(t *testing.T) {
	type args struct {
		task *models.Task
		id   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := models.GetTaskByID(tt.args.task, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	type args struct {
		task *models.Task
		id   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := models.UpdateTask(tt.args.task, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	type args struct {
		task *models.Task
		id   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := models.DeleteTask(tt.args.task, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
