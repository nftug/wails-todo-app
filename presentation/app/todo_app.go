package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	usecase "github.com/nftug/wails-todo-app/usecase/todo"
)

type TodoApp struct {
	ctx              context.Context
	createTodo       *usecase.CreateTodoUseCase
	updateTodo       *usecase.UpdateTodoUseCase
	updateTodoStatus *usecase.UpdateTodoStatusUseCase
	deleteTodo       *usecase.DeleteTodoUseCase
	getTodo          *usecase.GetTodoUseCase
	searchTodo       *usecase.GetTodoListUseCase
}

func NewTodoApp(
	createTodo *usecase.CreateTodoUseCase,
	updateTodo *usecase.UpdateTodoUseCase,
	updateTodoStatus *usecase.UpdateTodoStatusUseCase,
	deleteTodo *usecase.DeleteTodoUseCase,
	getTodo *usecase.GetTodoUseCase,
	searchTodo *usecase.GetTodoListUseCase) *TodoApp {
	return &TodoApp{nil, createTodo, updateTodo, updateTodoStatus, deleteTodo, getTodo, searchTodo}
}

func (a *TodoApp) OnDomReady(ctx context.Context) {
	a.ctx = ctx
}

func (a *TodoApp) Create(command todo.CreateCommand) (*interfaces.CreatedResponse, error) {
	return a.createTodo.Execute(a.ctx, command)
}

func (a *TodoApp) Update(id string, command todo.UpdateCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.updateTodo.Execute(a.ctx, parsedID, command)
}

func (a *TodoApp) UpdateStatus(id string, command todo.UpdateStatusCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.updateTodoStatus.Execute(a.ctx, parsedID, command)
}

func (a *TodoApp) Delete(id string) error {
	parsedID, _ := uuid.Parse(id)
	return a.deleteTodo.Execute(a.ctx, parsedID)
}

func (a *TodoApp) GetDetail(id string) (*todo.DetailResponse, error) {
	parsedID, _ := uuid.Parse(id)
	return a.getTodo.Execute(a.ctx, parsedID)
}

func (a *TodoApp) Search(query todo.Query) ([]*todo.ItemResponse, error) {
	return a.searchTodo.Execute(a.ctx, query)
}
