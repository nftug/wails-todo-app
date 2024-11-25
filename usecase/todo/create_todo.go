package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/shared/dtos"
	"github.com/samber/do"
)

type CreateTodoUseCase interface {
	Execute(ctx context.Context, command todo.CreateCommand) (*dtos.CreatedResponse, error)
}

type createTodoUseCase struct {
	repo todo.TodoRepository
}

func NewCreateTodoUseCase(i *do.Injector) (CreateTodoUseCase, error) {
	return &createTodoUseCase{do.MustInvoke[todo.TodoRepository](i)}, nil
}

func (u *createTodoUseCase) Execute(ctx context.Context, command todo.CreateCommand) (*dtos.CreatedResponse, error) {
	t, err := todo.NewTodo(command)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Save(ctx, t); err != nil {
		return nil, err
	}
	return dtos.NewCreatedResponse(t), nil
}
