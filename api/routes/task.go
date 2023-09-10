package routes

import (
	"github.com/SemenVologdin/flag.test.go/api/controllers"
	"github.com/SemenVologdin/flag.test.go/api/lib"
	"github.com/gin-gonic/gin"
	"log"
)

type TaskRoute struct {
	handler    lib.RequestHandler
	controller controllers.TaskController
}

func NewTaskRoute(handler lib.RequestHandler, controller controllers.TaskController) TaskRoute {
	return TaskRoute{
		handler:    handler,
		controller: controller,
	}
}

func (r *TaskRoute) SetUp() {
	log.Println("Установка маршрутов на страницу!")

	api := r.handler.Gin.Group("/api").Use(func(context *gin.Context) {
		log.Println("Запрос получен!")
	})
	{
		api.GET("/tasks", r.controller.GetTasks)
		api.GET("/task/:id", r.controller.GetTask)
		api.PUT("/task/", r.controller.CreateTask)
		api.DELETE("/task/:id", r.controller.DeleteTask)
		api.PATCH("/task/:id", r.controller.UpdateTask)
	}
}

func (r *TaskRoute) Run(env lib.Env) error {
	port := env.ServerPort
	if port == "" {
		port = "8080"
	}
	return r.handler.Gin.Run(":" + port)
}
