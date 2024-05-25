package todo

import "time"

type CreateCommand struct {
	Title         string       `json:"title"`
	Description   *string      `json:"description"`
	InitialStatus *StatusValue `json:"initialStatus"`
	DueDate       *time.Time   `json:"dueDate"`
}

type UpdateCommand struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
}

type UpdateStatusCommand struct {
	Status StatusValue `json:"status"`
}
