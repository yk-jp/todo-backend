package models

type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	StatusRefer uint   `json:"status_refer"`
	Status      string `json:"status"`
}
