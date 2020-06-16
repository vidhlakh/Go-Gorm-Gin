package main

import (
	"go-gorm-gin/config"
	"go-gorm-gin/models"
	"go-gorm-gin/routes"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		log.Fatal("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Task{})
	r := routes.SetupRouter()
	//running
	r.Run()
}
