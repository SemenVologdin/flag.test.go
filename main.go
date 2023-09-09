package main

import (
	"github.com/SemenVologdin/flag.test.go/api/controllers"
	"github.com/SemenVologdin/flag.test.go/api/lib"
	"github.com/SemenVologdin/flag.test.go/api/repository"
	"github.com/SemenVologdin/flag.test.go/api/routes"
	"github.com/SemenVologdin/flag.test.go/api/services"
	"log"
)

func main() {
	log.Println("Старт приложения!")
	env := lib.NewEnv()
	database := lib.NewDatabase(env)
	repo := repository.NewTaskRepository(database)
	service := services.NewTaskService(repo)
	controller := controllers.NewTaskController(service)
	handler := lib.NewRequestHandler()

	route := routes.NewTaskRoute(handler, controller)
	route.SetUp()
	err := route.Run(env)
	if err != nil {
		log.Fatalf("Произошла ошибка при старте приложения!")
	}

}
