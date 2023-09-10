package controllers

import (
	"encoding/json"
	"github.com/SemenVologdin/flag.test.go/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskController struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) TaskController {
	return TaskController{service: service}
}

func (c TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := c.service.GetTasks()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (c TaskController) GetTask(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Не передан Id задачи!"})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Id задачи должно быть типа integer!"})
		return
	}

	task, err := c.service.GetTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	jsonTask, err := json.Marshal(task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": jsonTask})
}

func (c TaskController) CreateTask(ctx *gin.Context) {

}

func (c TaskController) UpdateTask(ctx *gin.Context) {
}

func (c TaskController) DeleteTask(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Не передан Id задачи!"})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Id задачи должно быть типа integer!"})
		return
	}

	err = c.service.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task_id": id})
}
