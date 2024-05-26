package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
)

type CreateTodoUseCase struct {
	repo todo.TodoRepository
}

func NewCreateTodoUseCase(repo todo.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{repo}
}

func (u *CreateTodoUseCase) Execute(ctx context.Context, command todo.CreateCommand) (*interfaces.CreatedResponse, error) {
	t, err := todo.NewTodo(command)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Save(ctx, t); err != nil {
		return nil, err
	}
	return interfaces.NewCreatedResponse(t), nil
}
