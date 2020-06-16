// Package controllers has functions called from Rest API which internally works with models
package controllers

import (
	"go-gorm-gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllTasks ... Get all tasks
func GetAllTasks(c *gin.Context) {

	var task []models.Task
	err := models.GetAllTasks(&task)
	if err != nil {
		//c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, task)
	}
}

//CreateTask ... Create Task
func CreateTask(c *gin.Context) {
	var task models.Task

	c.ShouldBindJSON(&task)
	err := models.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error in POST:": err.Error()})
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

//GetTaskByID ... Get the task by id
func GetTaskByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	err := models.GetTaskByID(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error in GET:": err.Error()})
		log.Fatal(err.Error())
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
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
		c.JSON(http.StatusNotFound, gin.H{"error in PUT reg ID:": err.Error()})
		log.Fatal(err)
		//c.JSON(http.StatusNotFound, task)
	}
	c.BindJSON(&task)
	err = models.UpdateTask(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error in PUT:": err.Error()})
		log.Fatal(err)
		//c.AbortWithStatus(http.StatusNotFound)
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
		c.JSON(http.StatusNotFound, gin.H{"error in Delete:": err.Error()})
		log.Fatal(err)
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
