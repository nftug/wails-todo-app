package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
)

type GetTodoUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) (*todo.DetailResponse, error)
}

type getTodoUseCase struct {
	query todo.TodoQueryService
}

func NewGetTodoUseCase(query todo.TodoQueryService) GetTodoUseCase {
	return &getTodoUseCase{query}
}

func (u *getTodoUseCase) Execute(ctx context.Context, id uuid.UUID) (*todo.DetailResponse, error) {
	t, err := u.query.Find(ctx, id)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, interfaces.NewNotFoundError("todo")
	}
	return t, nil
}
