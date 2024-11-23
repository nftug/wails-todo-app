package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/do"
)

type UpdateTodoStatusUseCase interface {
	Execute(ctx context.Context, id uuid.UUID, command todo.UpdateStatusCommand) error
}

type updateTodoStatusUseCase struct {
	repo todo.TodoRepository
}

func NewUpdateTodoStatusUseCase(i *do.Injector) (UpdateTodoStatusUseCase, error) {
	return &updateTodoStatusUseCase{do.MustInvoke[todo.TodoRepository](i)}, nil
}

func (u *updateTodoStatusUseCase) Execute(ctx context.Context, id uuid.UUID, command todo.UpdateStatusCommand) error {
	t, err := u.repo.Find(ctx, id)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError()
	}

	if err := t.UpdateStatus(command); err != nil {
		return err
	}

	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}
	return nil
}
