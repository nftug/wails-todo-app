package todo

import (
	"time"

	"github.com/google/uuid"
)

type ItemResponse struct {
	ID      uuid.UUID  `json:"id"`
	Title   string     `json:"title"`
	Status  StatusItem `json:"status"`
	DueDate *time.Time `json:"dueDate"`
}

type DetailResponse struct {
	ID              uuid.UUID  `json:"id"`
	Title           string     `json:"title"`
	Description     *string    `json:"description"`
	Status          StatusItem `json:"status"`
	StatusUpdatedAt time.Time  `json:"statusUpdatedAt"`
	DueDate         *time.Time `json:"dueDate"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
