package models

import "github.com/jackc/pgx/pgtype"

type Task struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	CreatedDate pgtype.Date `json:"date_create"`
	UpdatedDate pgtype.Date `json:"date_update"`

	Answers []Answer `json:"answers"`
}

func (t Task) GetTableName() string {
	return "tasks"
}
