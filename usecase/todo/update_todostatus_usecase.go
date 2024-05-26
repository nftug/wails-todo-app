package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
)

type UpdateTodoStatusUseCase struct {
	repo todo.TodoRepository
}

func NewUpdateTodoStatusUseCase(repo todo.TodoRepository) *UpdateTodoStatusUseCase {
	return &UpdateTodoStatusUseCase{repo}
}

func (u *UpdateTodoStatusUseCase) Execute(ctx context.Context, id uuid.UUID, command todo.UpdateStatusCommand) error {
	t, err := u.repo.Find(ctx, id)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError("todo")
	}

	if err := t.UpdateStatus(command); err != nil {
		return err
	}

	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}
	return nil
}
