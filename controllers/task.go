// Package controllers has functions called from Rest API which internally works with models
package controllers

import (
	"fmt"
	"go-gorm-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllTasks ... Get all tasks
func GetAllTasks(c *gin.Context) {

	var user []models.Task
	err := models.GetAllTasks(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

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

//GetTaskByID ... Get the task by id
func GetTaskByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.Task
	err := models.GetTaskByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// //GetTaskByTaskstatus ... Get the task by status
// func GetTaskByTaskstatus(c *gin.Context) {
// 	status := c.Params.ByName("taskstatus")
// 	var task models.Task
// 	err := models.GetTaskByTaskstatus(&task, status)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, task)
// 	}
// }

//UpdateTask ... Update the task information
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Params.ByName("id")
	err := models.GetTaskByID(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, task)
	}
	c.BindJSON(&task)
	err = models.UpdateTask(&task, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

//DeleteTask ... Delete the task
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Params.ByName("id")
	err := models.DeleteTask(&task, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
