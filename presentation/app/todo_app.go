package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
	usecase "github.com/nftug/wails-todo-app/usecase/todo"
	"github.com/samber/do"
)

type TodoApp struct {
	ctx              context.Context
	createTodo       usecase.CreateTodoUseCase
	updateTodo       usecase.UpdateTodoUseCase
	updateTodoStatus usecase.UpdateTodoStatusUseCase
	deleteTodo       usecase.DeleteTodoUseCase
	getTodo          usecase.GetTodoUseCase
	searchTodo       usecase.GetTodoListUseCase
	notifyTodo       usecase.NotifyTodoUseCase
}

func NewTodoApp(i *do.Injector) (*TodoApp, error) {
	return &TodoApp{
		nil,
		do.MustInvoke[usecase.CreateTodoUseCase](i),
		do.MustInvoke[usecase.UpdateTodoUseCase](i),
		do.MustInvoke[usecase.UpdateTodoStatusUseCase](i),
		do.MustInvoke[usecase.DeleteTodoUseCase](i),
		do.MustInvoke[usecase.GetTodoUseCase](i),
		do.MustInvoke[usecase.GetTodoListUseCase](i),
		do.MustInvoke[usecase.NotifyTodoUseCase](i),
	}, nil
}

func (a *TodoApp) OnDomReady(ctx context.Context) {
	a.ctx = ctx

	a.notifyTodo.Execute(ctx)
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

func (a *TodoApp) GetDetails(id string) (*todo.DetailResponse, error) {
	parsedID, _ := uuid.Parse(id)
	return a.getTodo.Execute(a.ctx, parsedID)
}

func (a *TodoApp) Search(query todo.Query) ([]*todo.ItemResponse, error) {
	return a.searchTodo.Execute(a.ctx, query)
}
