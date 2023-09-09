package services

import "github.com/SemenVologdin/flag.test.go/api/repository"

type TaskService struct {
	repository repository.TaskRepository
}

func NewTaskService(repository repository.TaskRepository) TaskService {
	return TaskService{repository: repository}
}

func (service TaskService) GetTasks() []interface{} {
	rows, _ := service.repository.DB.Query("SELECT * FROM pg_catalog.pg_tables")
	val, _ := rows.Values()

	return val
}
