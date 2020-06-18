package models

import (
	"go-gorm-gin/config"
	"go-gorm-gin/models"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

var deadline, _ = time.Parse(time.RFC3339, "2020-06-18T13:13:23.63608771-04:00")

type tests struct {
	ID          int
	Name        string
	Description string
	Deadline    time.Time
	Taskstatus  string
}

var testTask = []tests{
	{
		ID:          1,
		Name:        "Register",
		Description: "Register for conference",
		Deadline:    deadline,
		Taskstatus:  "Active",
	},
}

func TestGetAllTasks(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "description", "deadline", "taskstatus"}).
		AddRow(1, "Register", "Register for conference", "2020-06-18T13:13:23.63608771-04:00", "Active")
		//AddRow(1, "Renew", "Renew the license", "2020-07-8T13:13:23.63608771-03:00", "Active")

	mock.ExpectQuery(`SELECT * FROM "tasks"`).WillReturnRows(rows)

	var testTask *[]models.Task
	if err = config.DB.Find(testTask).Error; err != nil {
		t.Error(err.Error())
	}
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expections: %s", err)
	// }
}

func TestCreateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "description", "deadline", "taskstatus"}).
		AddRow(1, "Register", "Register for conference", "2020-06-18T13:13:23.63608771-04:00", "Active")
		//AddRow(1, "Renew", "Renew the license", "2020-07-8T13:13:23.63608771-03:00", "Active")

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "TodolistDB"`)).WithArgs(rows)

	// if err = config.DB.Create(testTask).Error; err != nil {
	// 	t.Errorf("Error:", err)
	// }
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
