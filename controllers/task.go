// Package controllers has functions called from Rest API which internally works with models
package controllers

import (
	"fmt"
	"go-gorm-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	err := models.CreateUser(&task)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}
