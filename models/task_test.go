package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func InitMockDB() {

}
func TestGetAllTasks(t *testing.T) {
	type args struct {
		task *[]Task
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
			if err := GetAllTasks(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	type args struct {
		task *Task
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
			if err := CreateTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	type args struct {
		task *Task
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
			if err := GetTaskByID(tt.args.task, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	type args struct {
		task *Task
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
			if err := UpdateTask(tt.args.task, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	type args struct {
		task *Task
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
			if err := DeleteTask(tt.args.task, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
