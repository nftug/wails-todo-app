package infrastructure

import (
	"github.com/nftug/wails-todo-app/infrastructure/common"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	injectCore(i)
	do.Provide(i, common.NewConfigPathService)
	do.Provide(i, todo.NewTodoNotificationSender)
}

func injectCore(i *do.Injector) {
	do.Provide(i, common.NewBBoltDB)
	do.Provide(i, todo.NewTodoRepository)
	do.Provide(i, todo.NewTodoQueryService)
}
