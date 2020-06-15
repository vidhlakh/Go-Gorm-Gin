// Config/Database.go

//Package config contains configuration information for the database Mysql
package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB create a pointer type gorm.db
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig is the function that returns DBConfig struct with values host,port, user, password, dbname
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "vidhya",
		Password: "password",
		DBName:   "TodolistDB",
	}
	return &dbConfig
}

// DbURL gets the DBConfig struct and returns in string format
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
