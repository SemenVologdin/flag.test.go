package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/SemenVologdin/flag.test.go/api/lib"
	"github.com/SemenVologdin/flag.test.go/api/models"
)

type TaskRepository struct {
	DB lib.DataBase
}

func NewTaskRepository(db lib.DataBase) TaskRepository {
	return TaskRepository{DB: db}
}

func (repo TaskRepository) GetTasks() ([]models.Task, error) {
	q, args, _ := sq.Select(
		"id",
		"title",
		"description",
		"created_at",
		"updated_at",
	).From("tasks").ToSql()

	rows, err := repo.DB.Query(q, args...)
	if err != nil {
		return make([]models.Task, 0), err
	}

	tasks := make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		if err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedDate, &task.UpdatedDate); err != nil {
			return make([]models.Task, 0), err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo TaskRepository) GetTask(id int) (models.Task, error) {
	q, args, err := sq.Select(
		"id",
		"title",
		"description",
		"created_at",
		"updated_at",
	).From("tasks").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return models.Task{}, err
	}

	rows, err := repo.DB.Query(q, args...)
	if err != nil {
		return models.Task{}, err
	}

	var task models.Task
	if err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedDate, &task.UpdatedDate); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (repo TaskRepository) DeleteTask(id int) error {
	q, args, err := sq.Delete("tasks").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.Query(q, args...)
	if err != nil {
		return err
	}

	return nil
}
