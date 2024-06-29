package infrastructure

import (
	"github.com/nftug/wails-todo-app/infrastructure/config"
	"github.com/nftug/wails-todo-app/infrastructure/persistence"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	injectCore(i)
	do.Provide(i, config.NewLocalPathService)
	do.Provide(i, persistence.NewDB)
}

func injectCore(i *do.Injector) {
	do.Provide(i, todo.NewTodoRepository)
	do.Provide(i, todo.NewTodoQueryService)
}
