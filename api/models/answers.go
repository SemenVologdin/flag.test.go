package models

type Answer struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	CreatedDate string `json:"date_create"`
	UpdatedDate string `json:"date_update"`
	TaskId      string `json:"task_id"`
}

func (a Answer) GetTableName() string {
	return "answers"
}
