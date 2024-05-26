package usecase

import (
	"github.com/nftug/wails-todo-app/usecase/todo"
	"gorm.io/gorm"
)

type UseCaseAdapter struct {
	CreateTodo       *todo.CreateTodoUseCase
	UpdateTodo       *todo.UpdateTodoUseCase
	UpdateTodoStatus *todo.UpdateTodoStatusUseCase
	DeleteTodo       *todo.DeleteTodoUseCase
	GetTodo          *todo.GetTodoUseCase
	SearchTodo       *todo.GetTodoListUseCase
}

func NewUseCaseAdapter(
	createTodo *todo.CreateTodoUseCase,
	updateTodo *todo.UpdateTodoUseCase,
	updateTodoStatus *todo.UpdateTodoStatusUseCase,
	deleteTodo *todo.DeleteTodoUseCase,
	getTodo *todo.GetTodoUseCase,
	searchTodo *todo.GetTodoListUseCase) *UseCaseAdapter {
	return &UseCaseAdapter{
		createTodo, updateTodo, updateTodoStatus, deleteTodo, getTodo, searchTodo,
	}
}

type UseCaseAdapterMock struct {
	DB *gorm.DB
	UseCaseAdapter
}

func NewUseCaseAdapterMock(db *gorm.DB, u *UseCaseAdapter) *UseCaseAdapterMock {
	return &UseCaseAdapterMock{db, *u}
}
