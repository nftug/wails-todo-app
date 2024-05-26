package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
)

type UpdateTodoUseCase struct {
	repo todo.TodoRepository
}

func NewUpdateTodoUseCase(repo todo.TodoRepository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{repo}
}

func (u *UpdateTodoUseCase) Execute(ctx context.Context, id uuid.UUID, command todo.UpdateCommand) error {
	t, err := u.repo.Find(ctx, id)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError("todo")
	}

	if err := t.Update(command); err != nil {
		return err
	}

	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}
	return nil
}
