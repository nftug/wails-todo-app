package todo

import (
	"time"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/domain/todo/enums"
)

type TodoDBSchema struct {
	PK              int       `gorm:"primaryKey"`
	ID              uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	Title           string
	Description     *string
	Status          enums.StatusValue
	StatusUpdatedAt time.Time
	DueDate         *time.Time `gorm:"type:TIMESTAMP;null;default:null"`
	CreatedAt       time.Time  `gorm:"type:TIMESTAMP;null;"`
	UpdatedAt       *time.Time `gorm:"type:TIMESTAMP;null;default:null"`
}

func (t TodoDBSchema) TableName() string { return "todos" }

func (t *TodoDBSchema) ToEntity() *todo.Todo {
	return todo.Reconstruct(
		t.PK, t.ID,
		t.Title, t.Description,
		t.Status, t.StatusUpdatedAt,
		t.DueDate,
		t.CreatedAt, t.UpdatedAt,
	)
}

func (t *TodoDBSchema) Transfer(e *todo.Todo) *TodoDBSchema {
	return &TodoDBSchema{
		PK:              e.PK(),
		ID:              e.ID(),
		Title:           e.Title(),
		Description:     e.Description(),
		Status:          e.Status(),
		StatusUpdatedAt: e.StatusUpdatedAt(),
		DueDate:         e.DueDate(),
		CreatedAt:       e.CreatedAt(),
		UpdatedAt:       e.UpdatedAt(),
	}
}

func (t *TodoDBSchema) GetPK() int { return t.PK }

func (t *TodoDBSchema) ToDetailResponse() *todo.DetailResponse {
	return &todo.DetailResponse{
		ID:              t.ID,
		Title:           t.Title,
		Description:     t.Description,
		Status:          t.Status,
		StatusUpdatedAt: t.StatusUpdatedAt,
		DueDate:         t.DueDate,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
}

func (t *TodoDBSchema) ToItemResponse() *todo.ItemResponse {
	return &todo.ItemResponse{
		ID:      t.ID,
		Title:   t.Title,
		Status:  t.Status,
		DueDate: t.DueDate,
	}
}
