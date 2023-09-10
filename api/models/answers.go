package models

import "github.com/jackc/pgx/pgtype"

type Answer struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	CreatedDate pgtype.Date `json:"date_create"`
	UpdatedDate pgtype.Date `json:"date_update"`
	IsCorrect   bool        `json:"is_correct"`
	TaskId      int         `json:"task_id"`
}

func (a Answer) GetTableName() string {
	return "answers"
}
