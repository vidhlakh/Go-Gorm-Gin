package models

import (
	"fmt"
	"go-gorm-gin/config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAlltasks Fetch all task data
func GetAllTasks(task *[]Task) (err error) {
	if err = config.DB.Find(task).Error; err != nil {
		return err
	}
	return nil
}

//CreateTask ... Insert New data
func CreateTask(task *Task) (err error) {
	if err = config.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}

//GetTaskByID ... Fetch only one task by Id
func GetTaskByID(task *Task, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(task).Error; err != nil {
		return err
	}
	return nil
}

// //GetTaskByTaskstatus ... Fetch only one task by Status
// func GetTaskByTaskstatus(task *Task, taskstatus string) (err error) {
// 	if err = config.DB.Where("taskstatus = ?", taskstatus).Find(task).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

//UpdateTask ... Update task
func UpdateTask(task *Task, id string) (err error) {
	fmt.Println(task)
	config.DB.Save(task)
	return nil
}

//DeleteTask ... Delete task
func DeleteTask(task *Task, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(task)
	return nil
}
