package models

import "go-gorm-gin/config"

//CreateUser ... Insert New data
func CreateUser(task *Task) (err error) {
	if err = config.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}
