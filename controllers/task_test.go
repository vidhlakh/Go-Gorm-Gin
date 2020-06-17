// Package controllers has functions called from Rest API which internally works with models
package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/todolist/task")
	req, err := http.NewRequest(http.MethodGet, "/todolist/task", nil)
	if err != nil {
		t.Fatalf("Couldn't create request %v\n", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, resp.Code)
	}

}

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/todolist/task")

	req, err := http.NewRequest(http.MethodPost, "/todolist/task", nil)
	if err != nil {
		t.Fatalf("Couldn't create request %v\n", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code == http.StatusCreated {
		t.Log("Successful create ")
	} else if resp.Code == http.StatusBadRequest {
		t.Log("Couldn't create Request. Improper Request body ")
	} else if resp.Code == http.StatusOK {
		t.Log("Successful create")
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusCreated, resp.Code)
	}
}

func TestGetTaskByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/todolist/task/:id")
	req, err := http.NewRequest(http.MethodGet, "/todolist/task/:id", nil)
	if err != nil {
		t.Fatalf("Couldn't get Task for the id %v\n", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, resp.Code)
	}

}

func TestUpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/todolist/task/:id")
	req, err := http.NewRequest(http.MethodPut, "/todolist/task/:id", strings.NewReader("z=post"))
	if err != nil {
		t.Fatalf("Couldn't update request %v\n", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code == http.StatusAccepted {
		t.Log("Successful update ")
	} else if resp.Code == http.StatusBadRequest {
		t.Log("Couldn't update Request. Improper Request body ")
	} else if resp.Code == http.StatusOK {
		t.Log("Successful update")
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusAccepted, resp.Code)
	}

}

func TestDeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/todolist/task/:id")
	req, err := http.NewRequest(http.MethodDelete, "/todolist/task/:id", strings.NewReader("z=post"))
	if err != nil {
		t.Fatalf("Couldn't delete request %v\n", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, resp.Code)
	}

}
