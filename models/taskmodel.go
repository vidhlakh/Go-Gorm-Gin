//Models/Usermodel.go

// Package models holds the struct params for the table
package models

import "github.com/jinzhu/gorm"

// Task struct is the object to be passed as JSON to finally get stored in user table
type Task struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// // TableName changes the default table name to user
// func (b *Task) TableName() string {
// 	return "task"
// }
