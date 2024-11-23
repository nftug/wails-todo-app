package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/do"
)

type UpdateTodoUseCase interface {
	Execute(ctx context.Context, id uuid.UUID, command todo.UpdateCommand) error
}

type updateTodoUseCase struct {
	repo todo.TodoRepository
}

func NewUpdateTodoUseCase(i *do.Injector) (UpdateTodoUseCase, error) {
	return &updateTodoUseCase{do.MustInvoke[todo.TodoRepository](i)}, nil
}

func (u *updateTodoUseCase) Execute(ctx context.Context, id uuid.UUID, command todo.UpdateCommand) error {
	t, err := u.repo.Find(ctx, id)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError()
	}

	if err := t.Update(command); err != nil {
		return err
	}

	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}
	return nil
}
