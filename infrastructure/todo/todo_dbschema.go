package todo

import (
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/shared/enums"
)

type TodoDBSchema struct {
	ID              int               `json:"id"`
	Title           string            `json:"title"`
	Description     *string           `json:"description"`
	Status          enums.StatusValue `json:"status"`
	StatusUpdatedAt time.Time         `json:"statusUpdatedAt"`
	DueDate         *time.Time        `json:"dueDate"`
	NotifiedAt      *time.Time        `json:"notifiedAt"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       *time.Time        `json:"updatedAt"`
}

func (t TodoDBSchema) TableName() string { return "todos" }

func (t *TodoDBSchema) ToEntity() *todo.Todo {
	return todo.Reconstruct(
		t.ID,
		t.Title, t.Description,
		t.Status, t.StatusUpdatedAt,
		t.DueDate,
		t.NotifiedAt,
		t.CreatedAt, t.UpdatedAt,
	)
}

func (t *TodoDBSchema) Transfer(e *todo.Todo) *TodoDBSchema {
	return &TodoDBSchema{
		ID:              e.ID(),
		Title:           e.Title(),
		Description:     e.Description(),
		Status:          e.Status(),
		StatusUpdatedAt: e.StatusUpdatedAt(),
		DueDate:         e.DueDate(),
		NotifiedAt:      e.NotifiedAt(),
		CreatedAt:       e.CreatedAt(),
		UpdatedAt:       e.UpdatedAt(),
	}
}

func (t *TodoDBSchema) GetID() int { return t.ID }

func (t *TodoDBSchema) ToDetailsResponse() *todo.DetailsResponse {
	return &todo.DetailsResponse{
		ID:              t.ID,
		Title:           t.Title,
		Description:     t.Description,
		Status:          t.Status,
		StatusUpdatedAt: t.StatusUpdatedAt,
		DueDate:         t.DueDate,
		NotifiedAt:      t.NotifiedAt,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
}

func (t *TodoDBSchema) ToItemResponse() *todo.ItemResponse {
	return &todo.ItemResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		DueDate:     t.DueDate,
		NotifiedAt:  t.NotifiedAt,
	}
}
