package controllers

import (
	"github.com/SemenVologdin/flag.test.go/api/models"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не передан Id задачи!"})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id задачи должно быть типа integer!"})
		return
	}

	task, err := c.service.GetTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": task})
}

func (c TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskId, err := c.service.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": taskId})
}

func (c TaskController) UpdateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskId, err := c.service.UpdateTask(task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": taskId})
}

func (c TaskController) DeleteTask(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не передан Id задачи!"})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id задачи должно быть типа integer!"})
		return
	}

	err = c.service.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task_id": id})
}
