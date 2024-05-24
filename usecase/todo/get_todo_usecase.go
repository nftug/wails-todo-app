package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/shared/myerror"
	"github.com/nftug/wails-todo-app/domain/todo"
)

type GetTodoUseCase struct {
	query todo.TodoQueryService
}

func NewGetTodoUseCase(query todo.TodoQueryService) *GetTodoUseCase {
	return &GetTodoUseCase{query}
}

func (u *GetTodoUseCase) Execute(id uuid.UUID, ctx context.Context) (*todo.DetailResponse, error) {
	t, err := u.query.Find(id, ctx)
	if err != nil {
		return nil, err
	} else if t == nil {
		return nil, myerror.NewNotFoundError("todo")
	}
	return t, nil
}
