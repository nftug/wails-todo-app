package todo

import (
	"time"

	"github.com/nftug/wails-todo-app/interfaces/enums"
)

type ItemResponse struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description *string           `json:"description"`
	Status      enums.StatusValue `json:"status"`
	NotifiedAt  *time.Time        `json:"notifiedAt"`
	DueDate     *time.Time        `json:"dueDate"`
}

type DetailResponse struct {
	ID              string            `json:"id"`
	Title           string            `json:"title"`
	Description     *string           `json:"description"`
	Status          enums.StatusValue `json:"status"`
	StatusUpdatedAt time.Time         `json:"statusUpdatedAt"`
	DueDate         *time.Time        `json:"dueDate"`
	NotifiedAt      *time.Time        `json:"notifiedAt"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       *time.Time        `json:"updatedAt"`
}
