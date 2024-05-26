package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/nftug/wails-todo-app/usecase"
)

type TodoApp struct {
	ctx     context.Context
	adapter *usecase.UseCaseAdapter
}

func NewTodoApp(adapter *usecase.UseCaseAdapter) *TodoApp { return &TodoApp{nil, adapter} }

func (a *TodoApp) OnDomReady(ctx context.Context) {
	a.ctx = ctx
}

func (a *TodoApp) Create(command todo.CreateCommand) (*interfaces.CreatedResponse, error) {
	return a.adapter.CreateTodo.Execute(a.ctx, command)
}

func (a *TodoApp) Update(id string, command todo.UpdateCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.UpdateTodo.Execute(a.ctx, parsedID, command)
}

func (a *TodoApp) UpdateStatus(id string, command todo.UpdateStatusCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.UpdateTodoStatus.Execute(a.ctx, parsedID, command)
}

func (a *TodoApp) Delete(id string) error {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.DeleteTodo.Execute(a.ctx, parsedID)
}

func (a *TodoApp) GetDetail(id string) (*todo.DetailResponse, error) {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.GetTodo.Execute(a.ctx, parsedID)
}

func (a *TodoApp) Search(query todo.Query) ([]*todo.ItemResponse, error) {
	return a.adapter.SearchTodo.Execute(a.ctx, query)
}
