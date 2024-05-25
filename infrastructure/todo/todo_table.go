package todo

import (
	"time"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
)

type TodoTable struct {
	PK              int       `gorm:"primaryKey"`
	ID              uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	Title           string
	Description     *string
	Status          todo.StatusItem
	StatusUpdatedAt time.Time
	DueDate         *time.Time
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

func (t TodoTable) ToEntity() *todo.Todo {
	return todo.Reconstruct(
		t.PK, t.ID,
		t.Title, t.Description,
		t.Status, t.StatusUpdatedAt,
		t.DueDate,
		t.CreatedAt, t.UpdatedAt,
	)
}

func (t *TodoTable) Transfer(e *todo.Todo) {
	t.PK = e.PK()
	t.ID = e.ID()
	t.Title = e.Title()
	t.Description = e.Description()
	t.Status = e.Status()
	t.StatusUpdatedAt = e.StatusUpdatedAt()
	t.DueDate = e.DueDate()
	t.CreatedAt = e.CreatedAt()
	t.UpdatedAt = e.UpdatedAt()
}

func (t TodoTable) GetPK() int { return t.PK }

func (t TodoTable) ToDetailResponse() *todo.DetailResponse {
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

func (t TodoTable) ToItemResponse() *todo.ItemResponse {
	return &todo.ItemResponse{
		ID:      t.ID,
		Title:   t.Title,
		Status:  t.Status,
		DueDate: t.DueDate,
	}
}

func (t TodoTable) TableName() string { return "todos" }
