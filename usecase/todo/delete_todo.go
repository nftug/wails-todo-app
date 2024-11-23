package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/do"
)

type DeleteTodoUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) error
}

type deleteTodoUseCase struct {
	repo todo.TodoRepository
}

func NewDeleteTodoUseCase(i *do.Injector) (DeleteTodoUseCase, error) {
	return &deleteTodoUseCase{do.MustInvoke[todo.TodoRepository](i)}, nil
}

func (u *deleteTodoUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	t, err := u.repo.Find(ctx, id)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError()
	}

	if err := u.repo.Delete(ctx, t); err != nil {
		return err
	}
	return nil
}
