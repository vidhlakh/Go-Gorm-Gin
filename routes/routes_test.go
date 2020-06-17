// Package routes holds the endpoitns for the REST API
package routes

import (
	"go-gorm-gin/routes"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetupRouter(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// grp1 := r.Group("/todolist")
	// {
	// 	grp1.GET("task", controllers.GetAllTasks)
	// 	grp1.POST("task", controllers.CreateTask)
	// 	grp1.GET("task/:id", controllers.GetTaskByID)
	// 	//grp1.GET("task/:taskstatus", controllers.GetTaskByTaskstatus)
	// 	grp1.PUT("task/:id", controllers.UpdateTask)
	// 	grp1.DELETE("task/:id", controllers.DeleteTask)
	// }

	tests := []struct {
		name string
		want *gin.Engine
	}{
		{
			name: "Test Router",
			want: r,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := routes.SetupRouter()

			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("SetupRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
