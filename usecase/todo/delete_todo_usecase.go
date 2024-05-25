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

func (u *DeleteTodoUseCase) Execute(id uuid.UUID, ctx context.Context) error {
	t, err := u.repo.Find(id, ctx)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError("todo")
	}

	if err := u.repo.Delete(t, ctx); err != nil {
		return err
	}
	return nil
}
