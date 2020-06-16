// Package routes holds the endpoitns for the REST API
package routes

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetupRouter(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetupRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetupRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
