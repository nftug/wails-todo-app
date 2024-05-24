package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/usecase/shared/response"
	usecase "github.com/nftug/wails-todo-app/usecase/todo"
)

type TodoApp struct {
	create       *usecase.CreateTodoUseCase
	update       *usecase.UpdateTodoUseCase
	updateStatus *usecase.UpdateTodoStatusUseCase
	delete       *usecase.DeleteTodoUseCase
	get          *usecase.GetTodoUseCase
	search       *usecase.GetTodoListUseCase
}

func NewTodoApp(
	create *usecase.CreateTodoUseCase,
	update *usecase.UpdateTodoUseCase,
	updateStatus *usecase.UpdateTodoStatusUseCase,
	delete *usecase.DeleteTodoUseCase,
	get *usecase.GetTodoUseCase,
	search *usecase.GetTodoListUseCase) *TodoApp {
	return &TodoApp{create, update, updateStatus, delete, get, search}
}

func (a *TodoApp) Create(command todo.CreateCommand) (*response.CreatedResponse, error) {
	return a.create.Execute(command, context.Background())
}

func (a *TodoApp) Update(id string, command todo.UpdateCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.update.Execute(parsedID, command, context.Background())
}

func (a *TodoApp) UpdateStatus(id string, command todo.UpdateStatusCommand) error {
	parsedID, _ := uuid.Parse(id)
	return a.updateStatus.Execute(parsedID, command, context.Background())
}

func (a *TodoApp) Delete(id string) error {
	parsedID, _ := uuid.Parse(id)
	return a.delete.Execute(parsedID, context.Background())
}

func (a *TodoApp) GetDetail(id string) (*todo.DetailResponse, error) {
	parsedID, _ := uuid.Parse(id)
	return a.get.Execute(parsedID, context.Background())
}

func (a *TodoApp) Search(query todo.Query) ([]*todo.ItemResponse, error) {
	return a.search.Execute(query, context.Background())
}
