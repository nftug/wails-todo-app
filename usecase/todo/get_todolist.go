package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/do"
)

type GetTodoListUseCase interface {
	Execute(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error)
}

type getTodoListUseCase struct {
	query todo.TodoQueryService
}

func NewGetTodoListUseCase(i *do.Injector) (GetTodoListUseCase, error) {
	return &getTodoListUseCase{do.MustInvoke[todo.TodoQueryService](i)}, nil
}

func (u *getTodoListUseCase) Execute(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error) {
	t, err := u.query.FindAll(ctx, query)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, interfaces.NewNotFoundError()
	}
	return t, nil
}
