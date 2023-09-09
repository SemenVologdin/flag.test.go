package controllers

import (
	"github.com/SemenVologdin/flag.test.go/api/services"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service services.TaskService
}

func (c TaskController) GetTasks(ctx *gin.Context) {
}

func (c TaskController) GetTask(ctx *gin.Context) {
}

func (c TaskController) CreateTask(ctx *gin.Context) {
}

func (c TaskController) UpdateTask(ctx *gin.Context) {
}

func (c TaskController) DeleteTask(ctx *gin.Context) {
}

func NewTaskController(service services.TaskService) TaskController {
	return TaskController{service: service}
}
