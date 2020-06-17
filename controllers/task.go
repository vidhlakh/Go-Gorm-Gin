// Package controllers has functions called from Rest API which internally works with models
package controllers

import (
	"fmt"
	"go-gorm-gin/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//GetAllTasks ... Get all tasks
func GetAllTasks(c *gin.Context) {

	var task []models.Task

	err := models.GetAllTasks(&task)
	if err != nil {
		//c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	} else {
		c.JSON(http.StatusOK, task)
	}
}

//CreateTask ... Create Task
func CreateTask(c *gin.Context) {
	var task models.Task
	c.ShouldBindJSON(&task)

	// If deadline time format is wrong/ empty, consider the time.Now()
	value, present := c.GetPostForm("Deadline")
	if present {
		fmt.Println("Value:", value)
		deadline, err := time.Parse(time.RFC3339, value)
		if err != nil {
			fmt.Println("Error in deadline:", err)
			log.Print(err)

		}
		task.Deadline = deadline
	} else {
		task.Deadline = time.Now()
	}

	err1 := models.CreateTask(&task)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
		log.Print(err1.Error())

	} else {
		c.JSON(http.StatusCreated, task)
	}
}

//GetTaskByID ... Get the task by id
func GetTaskByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	err := models.GetTaskByID(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()}) //record not found
		log.Print(err.Error())

	} else {
		c.JSON(http.StatusOK, task)
	}
}

//UpdateTask ... Update the task information
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Params.ByName("id")
	err := models.GetTaskByID(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		log.Print(err.Error())
		//c.JSON(http.StatusNotFound, task)
	}
	c.BindJSON(&task)
	err = models.UpdateTask(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		log.Print(err.Error())
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusAccepted, task)
	}
}

//DeleteTask ... Delete the task
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Params.ByName("id")
	err := models.DeleteTask(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		log.Print(err.Error())
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
