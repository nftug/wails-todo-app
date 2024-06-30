package usecase

import (
	"github.com/nftug/wails-todo-app/usecase/todo"
	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	do.Provide(i, todo.NewCreateTodoUseCase)
	do.Provide(i, todo.NewUpdateTodoUseCase)
	do.Provide(i, todo.NewUpdateTodoStatusUseCase)
	do.Provide(i, todo.NewDeleteTodoUseCase)
	do.Provide(i, todo.NewGetTodoUseCase)
	do.Provide(i, todo.NewGetTodoListUseCase)
	do.Provide(i, todo.NewNotifyTodoUseCase)
}
