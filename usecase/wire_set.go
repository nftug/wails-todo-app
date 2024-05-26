package usecase

import (
	"github.com/google/wire"
	"github.com/nftug/wails-todo-app/usecase/todo"
)

var Set = wire.NewSet(
	NewUseCaseAdapter,
	NewUseCaseAdapterMock,
	todo.NewCreateTodoUseCase,
	todo.NewUpdateTodoUseCase,
	todo.NewUpdateTodoStatusUseCase,
	todo.NewDeleteTodoUseCase,
	todo.NewGetTodoUseCase,
	todo.NewGetTodoListUseCase,
)
