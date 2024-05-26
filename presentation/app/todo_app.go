package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/nftug/wails-todo-app/usecase"
)

type TodoApp struct{ adapter *usecase.UseCaseAdapter }

func NewTodoApp(adapter *usecase.UseCaseAdapter) *TodoApp { return &TodoApp{adapter} }

func (a *TodoApp) Create(command todo.CreateCommand) (*interfaces.CreatedResponse, error) {
	return a.adapter.CreateTodo.Execute(command, context.Background())
}

func (a *TodoApp) Update(id string, command todo.UpdateCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.UpdateTodo.Execute(parsedID, command, context.Background())
}

func (a *TodoApp) UpdateStatus(id string, command todo.UpdateStatusCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.UpdateTodoStatus.Execute(parsedID, command, context.Background())
}

func (a *TodoApp) Delete(id string) error {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.DeleteTodo.Execute(parsedID, context.Background())
}

func (a *TodoApp) GetDetail(id string) (*todo.DetailResponse, error) {
	parsedID, _ := uuid.Parse(id)
	return a.adapter.GetTodo.Execute(parsedID, context.Background())
}

func (a *TodoApp) Search(query todo.Query) ([]*todo.ItemResponse, error) {
	return a.adapter.SearchTodo.Execute(query, context.Background())
}
