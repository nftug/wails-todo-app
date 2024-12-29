package app

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/shared/dtos"
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

func (a *TodoApp) CreateTodo(command todo.CreateCommand) (*dtos.CreatedResponse, error) {
	return a.createTodo.Execute(a.ctx, command)
}

func (a *TodoApp) UpdateTodo(id int, command todo.UpdateCommand) error {
	return a.updateTodo.Execute(a.ctx, id, command)
}

func (a *TodoApp) UpdateTodoStatus(id int, command todo.UpdateStatusCommand) error {
	return a.updateTodoStatus.Execute(a.ctx, id, command)
}

func (a *TodoApp) DeleteTodo(id int) error {
	return a.deleteTodo.Execute(a.ctx, id)
}

func (a *TodoApp) GetTodoDetails(id int) (*todo.DetailsResponse, error) {
	return a.getTodo.Execute(a.ctx, id)
}

func (a *TodoApp) GetTodoList(query todo.Query) ([]*todo.ItemResponse, error) {
	return a.searchTodo.Execute(a.ctx, query)
}
