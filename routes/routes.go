// Package routes holds the endpoitns for the REST API
package routes

import (
	"go-gorm-gin/controllers"

	"github.com/gin-gonic/gin"
)
//Setup Test Router ... Configure routes
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	//grp1 := r.Group("/todolist")
	// {
	// 	grp1.GET("task", controllers.GetAllTasks)
	// 	grp1.POST("task", controllers.CreateTask)
	// 	grp1.GET("task/:id", controllers.GetTaskByID)
	// 	//grp1.GET("task/:taskstatus", controllers.GetTaskByTaskstatus)
	// 	grp1.PUT("task/:id", controllers.UpdateTask)
	// 	grp1.DELETE("task/:id", controllers.DeleteTask)
	// }
	return r



//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/todolist")
	{
		grp1.GET("task", controllers.GetAllTasks)
		grp1.POST("task", controllers.CreateTask)
		grp1.GET("task/:id", controllers.GetTaskByID)
		//grp1.GET("task/:taskstatus", controllers.GetTaskByTaskstatus)
		grp1.PUT("task/:id", controllers.UpdateTask)
		grp1.DELETE("task/:id", controllers.DeleteTask)
	}
	return r
}
