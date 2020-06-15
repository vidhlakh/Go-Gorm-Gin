// Package controllers has functions called from Rest API which internally works with models
package controllers

import (
	"fmt"
	"go-gorm-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateTask ... Create Task
func CreateTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	err := models.CreateTask(&task)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}
