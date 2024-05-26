package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
)

type DeleteTodoUseCase struct {
	repo todo.TodoRepository
}

func NewDeleteTodoUseCase(repo todo.TodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{repo}
}

func (u *DeleteTodoUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	t, err := u.repo.Find(ctx, id)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError("todo")
	}

	if err := u.repo.Delete(ctx, t); err != nil {
		return err
	}
	return nil
}
