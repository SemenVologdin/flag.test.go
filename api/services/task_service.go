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
	tx, err := serv.repository.DB.Begin()
	if err != nil {
		return err
	}

	err = serv.repository.DeleteAnswersByTaskId(id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = serv.repository.DeleteTask(id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (serv TaskService) CreateTask(task models.Task) (int, error) {
	tx, err := serv.repository.DB.Begin()
	if err != nil {
		return 0, err
	}

	taskId, err := serv.repository.CreateTask(task)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = serv.repository.CreateAnswers(taskId, task.Answers)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return taskId, nil
}

func (serv TaskService) UpdateTask(task models.Task) (int, error) {
	tx, err := serv.repository.DB.Begin()
	if err != nil {
		return 0, err
	}

	taskId, err := serv.repository.UpdateTask(task)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = serv.repository.UpdateAnswers(taskId, task.Answers)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return taskId, nil
}
