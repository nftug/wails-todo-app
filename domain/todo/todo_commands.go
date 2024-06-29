package todo

import (
	"time"

	"github.com/nftug/wails-todo-app/domain/todo/enums"
)

type CreateCommand struct {
	Title         string             `json:"title"`
	Description   *string            `json:"description"`
	InitialStatus *enums.StatusValue `json:"initialStatus"`
	DueDate       *time.Time         `json:"dueDate"`
}

type UpdateCommand struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
}

type UpdateStatusCommand struct {
	Status enums.StatusValue `json:"status"`
}
