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

func (u *UpdateTodoUseCase) Execute(id uuid.UUID, command todo.UpdateCommand, ctx context.Context) error {
	t, err := u.repo.Find(id, ctx)
	if err != nil {
		return err
	} else if t == nil {
		return interfaces.NewNotFoundError("todo")
	}

	if err := t.Update(command); err != nil {
		return err
	}

	if err := u.repo.Save(t, ctx); err != nil {
		return err
	}
	return nil
}
