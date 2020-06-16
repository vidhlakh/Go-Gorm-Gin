//Models/Usermodel.go

// Package models holds the struct params for the table
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task struct is the object to be passed as JSON to finally get stored in user table
type Task struct {
	gorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	Taskstatus  string    `json:"taskstatus" default:"Active"`
}

//Sample JSON message
// {
//     "name":"Bill payment",
//     "description":"Pay the EB bill ",
//     "deadline":"2020-06-15",
//     "taskstatus":"Active"
// }
