package main

import (
	"fmt"
	"go-gorm-gin/config"
	"go-gorm-gin/models"
	"go-gorm-gin/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Task{})
	r := routes.SetupRouter()
	//running
	r.Run()
}
