package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/shared/customerr"
	"github.com/samber/do"
)

type UpdateTodoStatusUseCase interface {
	Execute(ctx context.Context, id int, command todo.UpdateStatusCommand) error
}

type updateTodoStatusUseCase struct {
	repo todo.TodoRepository
}

func NewUpdateTodoStatusUseCase(i *do.Injector) (UpdateTodoStatusUseCase, error) {
	return &updateTodoStatusUseCase{do.MustInvoke[todo.TodoRepository](i)}, nil
}

func (u *updateTodoStatusUseCase) Execute(ctx context.Context, id int, command todo.UpdateStatusCommand) error {
	t, err := u.repo.Find(ctx, id)
	if err != nil {
		return err
	} else if t == nil {
		return customerr.NewNotFoundError()
	}

	if err := t.UpdateStatus(command); err != nil {
		return err
	}

	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}
	return nil
}
