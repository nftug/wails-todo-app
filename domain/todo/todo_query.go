package todo

import (
	"time"

	"github.com/nftug/wails-todo-app/interfaces/enums"
)

type Query struct {
	Search      *string            `json:"search"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Status      *enums.StatusValue `json:"status"`
	Until       *time.Time         `json:"until"`
	After       *time.Time         `json:"after"`
	IsNotified  *bool              `json:"isNotified"`
}
