package repository

import "github.com/SemenVologdin/flag.test.go/api/lib"

type TaskRepository struct {
	DB lib.DataBase
}

func NewTaskRepository(db lib.DataBase) TaskRepository {
	return TaskRepository{DB: db}
}
