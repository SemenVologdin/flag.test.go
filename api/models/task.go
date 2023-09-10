package models

import "github.com/jackc/pgx/pgtype"

type Task struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	CreatedDate pgtype.Date `json:"created_at"`
	UpdatedDate pgtype.Date `json:"updated_at"`

	Answers []Answer `json:"answers"`
}

func (t Task) GetTableName() string {
	return "tasks"
}
