package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/SemenVologdin/flag.test.go/api/lib"
	"github.com/SemenVologdin/flag.test.go/api/models"
	"time"
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
	).
		From("tasks").
		ToSql()

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

		task.Answers = repo.getAnswersByTaskId(task.Id)
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
	).
		From("tasks").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return models.Task{}, err
	}

	rows, err := repo.DB.Query(q, args...)
	if err != nil {
		return models.Task{}, err
	}

	var task models.Task
	if !rows.Next() {
		return models.Task{}, fmt.Errorf("не удалось найти задачу")
	}

	if err = rows.Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.CreatedDate,
		&task.UpdatedDate); err != nil {
		return models.Task{}, err
	}

	task.Answers = repo.getAnswersByTaskId(task.Id)
	return task, nil
}

func (repo TaskRepository) DeleteTask(id int) error {
	q, args, err := sq.Delete("tasks").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo TaskRepository) CreateTask(task models.Task) (int, error) {
	datetime := time.Now().Format(time.RFC3339)

	q, args, err := sq.Insert("tasks").
		Columns("title", "description", "created_at", "updated_at").
		Values(task.Title, task.Description, datetime, datetime).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, err
	}

	row := repo.DB.QueryRow(q, args...)
	var recordId int
	if err := row.Scan(&recordId); err != nil {
		return 0, err
	}

	return recordId, nil
}

func (repo TaskRepository) UpdateTask(task models.Task) (int, error) {
	datetime := time.Now().Format(time.RFC3339)

	q, args, err := sq.Update(task.GetTableName()).
		Where(sq.Eq{"id": task.Id}).
		Set("title", task.Title).
		Set("description", task.Description).
		Set("updated_at", datetime).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	_, err = repo.DB.Exec(q, args...)
	if err != nil {
		return 0, err
	}

	return task.Id, nil
}

func (repo TaskRepository) CreateAnswers(taskId int, answers []models.Answer) error {
	if taskId == 0 {
		return fmt.Errorf("не передан Id теста")
	}

	if len(answers) == 0 {
		return nil
	}

	datetime := time.Now().Format(time.RFC3339)
	builder := sq.Insert("answers").
		Columns("title", "task_id", "created_at", "updated_at")

	for _, answer := range answers {
		builder = builder.Values(answer.Title, taskId, datetime, datetime)
	}

	q, args, err := builder.
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo TaskRepository) UpdateAnswers(taskId int, answers []models.Answer) error {
	for _, answer := range answers {
		if answer.Id != 0 {
			if err := repo.updateAnswerByTaskId(answer); err != nil {
				return err
			}
			continue
		}

		if err := repo.addAnswerByTaskId(taskId, answer); err != nil {
			return err
		}
	}

	return nil
}

func (repo TaskRepository) DeleteAnswersByTaskId(taskId int) error {
	if taskId == 0 {
		return fmt.Errorf("не передан id теста")
	}

	q, args, err := sq.Delete("answers").
		Where(sq.Eq{"task_id": taskId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo TaskRepository) getAnswersByTaskId(taskId int) []models.Answer {
	q, args, _ := sq.Select(
		"id",
		"title",
		"created_at",
		"updated_at",
		"task_id",
		"is_correct",
	).
		From("answers").
		Where(sq.Eq{"task_id": taskId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	answers := make([]models.Answer, 0)
	answersRows, err := repo.DB.Query(q, args...)
	if err != nil {
		return answers
	}

	for answersRows.Next() {
		var answer models.Answer
		if err = answersRows.Scan(
			&answer.Id,
			&answer.Title,
			&answer.CreatedDate,
			&answer.UpdatedDate,
			&answer.TaskId,
			&answer.IsCorrect,
		); err != nil {

			continue
		}
		answers = append(answers, answer)
	}

	return answers
}

func (repo TaskRepository) updateAnswerByTaskId(answer models.Answer) error {
	datetime := time.Now().Format(time.RFC3339)
	q, args, err := sq.Update(answer.GetTableName()).
		Where(sq.Eq{"id": answer.Id}).
		Set("title", answer.Title).
		Set("updated_at", datetime).
		Set("is_correct", answer.IsCorrect).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo TaskRepository) addAnswerByTaskId(taskId int, answer models.Answer) error {
	datetime := time.Now().Format(time.RFC3339)

	q, args, err := sq.Insert(answer.GetTableName()).
		Columns("title", "date_create", "date_update", "is_correct", "task_id").
		Values(answer.Title, datetime, datetime, answer.IsCorrect, taskId).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}
