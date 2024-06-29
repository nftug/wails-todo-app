package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
)

type GetTodoListUseCase interface {
	Execute(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error)
}

type getTodoListUseCase struct {
	query todo.TodoQueryService
}

func NewGetTodoListUseCase(query todo.TodoQueryService) GetTodoListUseCase {
	return &getTodoListUseCase{query}
}

func (u *getTodoListUseCase) Execute(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error) {
	t, err := u.query.FindAll(ctx, query)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, interfaces.NewNotFoundError("todo")
	}
	return t, nil
}
