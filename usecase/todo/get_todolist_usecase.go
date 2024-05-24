package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/shared/myerror"
	"github.com/nftug/wails-todo-app/domain/todo"
)

type GetTodoListUseCase struct {
	query todo.TodoQueryService
}

func NewGetTodoListUseCase(query todo.TodoQueryService) *GetTodoListUseCase {
	return &GetTodoListUseCase{query}
}

func (u *GetTodoListUseCase) Execute(query todo.Query, ctx context.Context) ([]*todo.ItemResponse, error) {
	t, err := u.query.FindAll(query, ctx)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, myerror.NewNotFoundError("todo")
	}
	return t, nil
}
