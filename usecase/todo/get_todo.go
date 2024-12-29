package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/shared/customerr"
	"github.com/samber/do"
)

type GetTodoUseCase interface {
	Execute(ctx context.Context, id int) (*todo.DetailsResponse, error)
}

type getTodoUseCase struct {
	query todo.TodoQueryService
}

func NewGetTodoUseCase(i *do.Injector) (GetTodoUseCase, error) {
	return &getTodoUseCase{do.MustInvoke[todo.TodoQueryService](i)}, nil
}

func (u *getTodoUseCase) Execute(ctx context.Context, id int) (*todo.DetailsResponse, error) {
	t, err := u.query.Find(ctx, id)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, customerr.NewNotFoundError()
	}
	return t, nil
}
