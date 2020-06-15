package models

import "go-gorm-gin/config"

//CreateTask ... Insert New data
func CreateTask(task *Task) (err error) {
	if err = config.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}
