package services

import (
	"github.com/SemenVologdin/flag.test.go/api/models"
	"github.com/SemenVologdin/flag.test.go/api/repository"
)

type TaskService struct {
	repository repository.TaskRepository
}

func NewTaskService(repository repository.TaskRepository) TaskService {
	return TaskService{repository: repository}
}

func (serv TaskService) GetTasks() ([]models.Task, error) {
	return serv.repository.GetTasks()
}

func (serv TaskService) GetTask(id int) (models.Task, error) {
	return serv.repository.GetTask(id)
}

func (serv TaskService) DeleteTask(id int) error {
	return serv.repository.DeleteTask(id)
}
