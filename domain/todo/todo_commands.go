package todo

import "time"

type CreateCommand struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      StatusItem `json:"status"`
	DueDate     *time.Time `json:"dueDate"`
}

type UpdateCommand struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
}

type UpdateStatusCommand struct {
	Status StatusItem `json:"status"`
}
