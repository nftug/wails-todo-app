package infrastructure

import (
	"github.com/nftug/wails-todo-app/infrastructure/middleware"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	injectCore(i)
	do.Provide(i, middleware.NewLocalPathService)
	do.Provide(i, middleware.NewDB)
}

func injectCore(i *do.Injector) {
	do.Provide(i, todo.NewTodoRepository)
	do.Provide(i, todo.NewTodoQueryService)
}
